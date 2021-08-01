package sdk

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"
)

func Start(handler interface{}) {
	paramsString := flag.String("wfparams", "", "params")
	wfInvoker := flag.Bool("wfInvoker", false, "workflow is not invoker")
	flag.Parse()

	if !*wfInvoker {
		fmt.Println("workflow is not invoker")
		return
	}

	if reflect.TypeOf(handler).Kind() != reflect.Func {
		fmt.Println("handler is not func")
		return
	}
	var params interface{}

	if *paramsString != "" {
		// decode base64 string to byte array
		paramsBytes, err := base64.StdEncoding.DecodeString(*paramsString)
		if err != nil {
			fmt.Println("invalid input params, check workflow config")
			return
		}
		// convert byte array to interface
		err = json.Unmarshal(paramsBytes, &params)
		if err != nil {
			fmt.Println("something went wrong in starting, please try again")
			return
		}
	}
	inputData := map[string]interface{}{
		"data": params,
	}
	in := []reflect.Value{reflect.ValueOf(inputData)}
	values := reflect.ValueOf(handler).Call(in)
	for _, v := range values {
		data, _ := json.Marshal(v.Interface())
		//encode interface to base64 string
		sEnc := base64.StdEncoding.EncodeToString(data)
		fmt.Printf("[[##WORKFLOW-START##]]%s[[##WORKFLOW-END##]]\n", sEnc)
	}
}
