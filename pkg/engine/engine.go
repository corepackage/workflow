package engine

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/corepackage/workflow/internal/constants"
	"github.com/corepackage/workflow/pkg/util"
)

type Resp struct {
	stepID string
	resp   interface{}
}

// Running the workflow
func (wf *Workflow) Run(ctx context.Context, headers map[string][]string, queryParams map[string][]string, body interface{}) (interface{}, error) {
	// fmt.Println(wf.Steps[0])

	userCtx := make(map[string]interface{})
	userCtx["body"] = body
	// TODO: fetching instance from db
	var lastID string
	waitCh := make(chan struct{})
	ctxCh := make(chan struct{})
	errCh := make(chan error)
	respCh := make(chan Resp, 10)

	// writing resonse to userContext
	go func() {
		for v := range respCh {
			userCtx[v.stepID] = v.resp
		}
	}()

	for _, step := range wf.Steps {
		var err error
		lastID = step.ID

		// verifying step type
		var newStep ExecuteStep
		if step.Type == constants.API_STEP {
			newStep = step.APIStep
		} else if step.Type == constants.LOGIC_STEP {
			newStep = step.LogicStep
		} else {
			return nil, errors.New("invalid step stype")
		}

		var (
			stepCtx context.Context
			cancel  context.CancelFunc
		)

		// Preparing timeout if defined
		if step.Timeout != "" && !step.Async {
			var timeout time.Duration
			timeout, err = util.ToTime(step.Timeout)
			if err != nil {
				log.Printf("Invalid Timeout value : %v\n", step.Timeout)
				return nil, err
			}
			stepCtx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()

			go func() {
				var err error
				select {
				case <-waitCh:
					fmt.Println("step complete before context")
					err = nil
				case <-stepCtx.Done():
					fmt.Println("context time out")
					err = errors.New("step execution timeout")
				}
				fmt.Println("writing to ctx channel")
				ctxCh <- struct{}{}
				errCh <- err
			}()
		}

		// If step is synced and delay is specified, wait for delay
		if !step.Async {
			// Adding delay to execution
			var delay time.Duration
			if step.Delay != "" {
				delay, err = util.ToTime(step.Delay)
				if err != nil {
					log.Printf("Invalid delay value : %v\n", step.Delay)
					return nil, err
				}
			}
			if delay > 0 {
				log.Printf("Sleeping for %v\n", delay)
				time.Sleep(delay)
			}
		}

		// Executing the step
		go func() {
			resp, err := newStep.Execute(wf, headers, queryParams, userCtx)
			if !step.Async {
				waitCh <- struct{}{}
			}
			respCh <- Resp{step.ID, resp}
			errCh <- err
		}()

		// If step is synced then waiting for timeout or response
		if !step.Async {
			fmt.Println("waiting for step to complete")
			select {
			case <-ctxCh:
			case <-waitCh:
			}
			fmt.Println("sync step complete")
			fmt.Println("waiting for error")
			err = <-errCh
			fmt.Println("step completed")

			if err != nil {
				log.Println("Run : Error executing API step")
				return nil, err
			}
			if step.Break {
				return userCtx[step.ID], nil
			}
		}
		fmt.Println("async step")

	}
	return userCtx[lastID], nil

}
