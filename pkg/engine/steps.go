package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/corepackage/workflow/internal/constants"
	"github.com/corepackage/workflow/pkg/util"
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

		resp, err := mapParams(queryParams, api.Endpoint)
		if err != nil {
			log.Println("API Execute : error mapping params ", err)
			return nil, errors.New("invalid expression for query params in endpoint")
		}
		var ok bool
		endpoint, ok = resp.(string)
		if !ok {
			log.Println("expected string after mapping")
			return nil, errors.New("API Execute : unexpected error")
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

// mapParams : mapping query params to iterface
func mapParams(queryParams map[string][]string, mapObj interface{}) (interface{}, error) {
	var strObj string
	switch v := mapObj.(type) {
	// If the obj is type string replace the string values
	case string:
		strObj = v
		matchStr := util.FindMatchStr(constants.QUERY_REGEX, strObj)
		for _, match := range matchStr {
			keys := strings.Split(match, ".")[1:]
			if len(keys) <= 0 {
				return nil, errors.New("cannot bind JSON object to string")
			}
			if len(keys) > 2 {
				return nil, errors.New("invalid pattern for query  params")
			}
			var key = keys[0]
			var strIndex string
			if len(keys) == 2 {
				strIndex = keys[1]
			} else {
				strIndex = "0"
			}

			val, ok := queryParams[key]
			if !ok {
				return nil, errors.New("invalid query params")
			}
			index, err := strconv.Atoi(strIndex)
			if err != nil {
				return nil, errors.New("query param index can only be integer")
			}
			if index < 0 || index >= len(val) {
				return nil, errors.New("query param index out of range")
			}
			strObj = strings.Replace(strObj, match, val[index], 1)
		}
		return strObj, nil
	// If the obj is type map replace the map values with whole object or replace recurrsively for each key in map
	case map[string]interface{}:
		for key, val := range v {
			str, ok := val.(string)
			if ok && str == "$$queryParams" {
				v[key] = queryParams
			} else {
				resp, err := mapParams(queryParams, val)
				if err != nil {
					return nil, errors.New("invalid query param")
				}
				v[key] = resp
			}
		}
		return v, nil

	default:
		return nil, errors.New("invalid obj type")
	}
}

// mapBody : mapping body to iterface
func mapBody(bodyJSON map[string]interface{}, mapObj interface{}) (interface{}, error) {
	var strObj string
	switch v := mapObj.(type) {
	// If the obj is type string replace the string values
	case string:
		strObj = v
		matchStr := util.FindMatchStr(constants.BODY_REGEX, strObj)
		var inputValue interface{}
		for _, match := range matchStr {
			keys := strings.Split(match, ".")[1:]
			if len(keys) <= 0 {
				return nil, errors.New("cannot bind JSON object to string")
			}
			var err error
			inputValue, err = findValue(bodyJSON, keys)
			if err != nil {
				log.Println("mapBody err", err)
				return nil, errors.New("key not found in input body for " + match)
			}
			res, ok := inputValue.(string)
			if ok {
				strObj = strings.Replace(strObj, match, res, 1)
				inputValue = strObj
			}
		}
		return inputValue, nil
	// If the obj is type map replace the map values with whole object or replace recurrsively for each key in map
	case map[string]interface{}:
		for key, val := range v {
			str, ok := val.(string)
			if ok && str == "$$body" {
				v[key] = bodyJSON
			} else {
				resp, err := mapBody(bodyJSON, val)
				if err != nil {
					fmt.Println(err)
					return nil, errors.New("invalid body key")
				}
				v[key] = resp
			}
		}
		return v, nil
	default:
		return nil, errors.New("invalid obj type")
	}
}
func findValue(bodyJson map[string]interface{}, keys []string) (interface{}, error) {
	itrMap := bodyJson
	var resp interface{}
	for i := 0; i < len(keys)-1; i++ {
		if itrMap[keys[i]] != nil {
			var ok bool
			itrMap, ok = itrMap[keys[i]].(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("key not found")
			}
		} else {
			return nil, fmt.Errorf("Key %s not found", keys[i])
		}
	}
	resp = itrMap[keys[len(keys)-1]]
	return resp, nil

}
