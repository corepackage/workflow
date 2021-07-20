package engine

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/corepackage/workflow/internal/constants"
)

// LogicStep - properties explicit to logic type step
type LogicStep struct {
	Runtime constants.Runtime `json:"runtime" yaml:"runtime"`
	ExePath string            `json:"exe-path" yaml:"exe-path"`
	Handler string            `json:"handler" yaml:"handler"`
}

// APIStep - properties explicit to logic type step
type APIStep struct {
	Endpoint       string            `json:"endpoint" yaml:"endpoint"`
	Method         string            `json:"method" yaml:"method"`
	Payload        interface{}       `json:"payload" yaml:"payload"`
	IncludeHeaders bool              `json:"include-headers" yaml:"include-headers"`
	CustomHeaders  map[string]string `json:"custom-headers" yaml:"custom-headers"`
}

// Execute : executing the logic function
func (l *LogicStep) Execute() {
}

func (api *APIStep) Execute(wf *Workflow, headers map[string][]string, queryParams map[string][]string, body map[string]interface{}) (interface{}, error) {
	var result interface{}
	var endpoint = api.Endpoint

	// Making http request for get
	if api.Method == http.MethodGet {

		// Replacing query params
		if strings.Contains(endpoint, "$$queryParams") {
			for k, v := range queryParams {
				endpoint = strings.Replace(endpoint, "$$queryParams."+k, v[0], -1)
			}
		}
		if strings.Contains(endpoint, "$$queryParams") {
			log.Println("API Execute error, query param not provided")
			return nil, errors.New("query param not provided")
		}

		if strings.Contains(endpoint, "$$body") {
			for k, v := range body {
				str, ok := v.(string)
				if !ok {
					log.Println("API Execute error, invalid key")
					return nil, errors.New("invalid key")
				}
				endpoint = strings.Replace(endpoint, "$$body."+k, str, -1)
			}
		}
		if strings.Contains(endpoint, "$$body") {
			log.Println("API Execute error, key not provided in body ")
			return nil, errors.New("key not provided in body ")
		}

	}
	//get request http
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		log.Println("API Step Execute err :", err)
		return nil, err
	}

	// add headers
	if api.IncludeHeaders {
		for k, v := range headers {
			req.Header[k] = v
		}
	}
	// Adding custom headers
	if len(api.CustomHeaders) != 0 {
		for k, v := range api.CustomHeaders {
			req.Header[k] = []string{v}
		}
	}
	// execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("API Step response create err ", err)
		return nil, err
	}
	defer resp.Body.Close()
	// read the response body
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("API Step response read err ", err)
		return nil, err
	}
	err = json.Unmarshal(byteArray, &result)
	if err != nil {
		log.Println("API Step unmarshal err ", err)
		return nil, err
	}
	return result, nil
}
