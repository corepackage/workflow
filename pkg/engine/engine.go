package engine

import (
	"log"

	"github.com/corepackage/workflow/internal/constants"
)

// Running the workflow
func (wf *Workflow) Run(headers map[string][]string, queryParams map[string][]string, body map[string]interface{}) (interface{}, error) {
	// fmt.Println(wf.Steps[0])

	resp := make(map[string]interface{})
	// TODO: fetching instance from db
	for _, step := range wf.Steps {
		var err error
		if step.Type == constants.API_STEP {
			resp[step.ID], err = step.APIStep.Execute(wf, headers, queryParams, body)
		}
		if err != nil {
			log.Println("Run : Error executing API step")
			return nil, err
		}
		if step.Break {
			return resp[step.ID], nil
		}
	}
	return nil, nil

}
