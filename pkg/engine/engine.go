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
	for _, step := range wf.Steps {
		var err error
		lastID = step.ID

		var newStep ExecuteStep
		if step.Type == constants.API_STEP {
			step.APIStep.payload = step.Payload
			newStep = step.APIStep
		} else if step.Type == constants.LOGIC_STEP {
			step.LogicStep.payload = step.Payload
			newStep = step.LogicStep
		} else {
			return nil, errors.New("invalid step stype")
		}

		var (
			stepCtx context.Context
			cancel  context.CancelFunc
		)
		if step.Timeout != "" {
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

		go func() {
			userCtx[step.ID], err = newStep.Execute(wf, headers, queryParams, userCtx)
			if err != nil {
				log.Println("Run : Error executing API step")
				errCh <- err
			}
			waitCh <- struct{}{}
		}()

		if !step.Async {
			fmt.Println("waiting for step to complete")
			select {
			case <-ctxCh:
			case <-waitCh:
			}
			fmt.Println("sync step complete")

		} else {
			fmt.Println("async step")
			<-ctxCh
		}

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
	return userCtx[lastID], nil

}
