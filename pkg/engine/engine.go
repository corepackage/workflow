package engine

import (
	"log"

	"github.com/corepackage/workflow/internal/constants"
)

// Running the workflow
func (wf *Workflow) Run(headers map[string][]string, queryParams map[string][]string, body interface{}) (interface{}, error) {
	// fmt.Println(wf.Steps[0])

	context := make(map[string]interface{})
	context["body"] = body
	context["testStep"] = map[string]interface{}{"ref": map[string]interface{}{"key": "from_test_step"}}
	// TODO: fetching instance from db
	for _, step := range wf.Steps {
		var err error
		if step.Type == constants.API_STEP {
			context[step.ID], err = step.APIStep.Execute(wf, headers, queryParams, context)
		}
		if err != nil {
			log.Println("Run : Error executing API step")
			return nil, err
		}
		if step.Break {
			return context[step.ID], nil
		}
	}
	return nil, nil

}
