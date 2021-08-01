package parser

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/corepackage/workflow/pkg/engine"
	"gopkg.in/yaml.v2"
)

var FilePath = RootDir()

// func main() {

// 	configFilePath := "/home/admino/Workspace/POCs/workflow_config.yml"
// 	err := FileYamlUnmarshal(configFilePath)
// 	if err != nil {
// 		log.Fatalf("Aborting : %v", err)
// 	}

// 	fmt.Println("Primay Keys : ", t.PrimaryKey[0].PKey)
// 	fmt.Println("Workflow Name : ", t.Name)
// 	fmt.Println("Cors : ", t.CORS)
// 	// fmt.Println(t.Steps[1].ID)
// 	for key, value := range t.Steps {
// 		fmt.Println("Key :", key, "Value :", value)
// 		fmt.Println("Step Name:", value.Name)
// 		fmt.Println("Pre Condition:", value.PreCondition)
// 		fmt.Println("Pre Condition:", value.PostCondition)
// 	}

// 	workflowId, err := GetID(configFilePath)
// 	fmt.Println(workflowId, err)
// 	workflowName, err := GetName(configFilePath)
// 	fmt.Println(workflowName, err)

// }

// RootDir : To get the root directory of the workflow project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../")
	return filepath.Dir(d)
}

// FileYamlUnmarshal : To unmarshal the YAML file to the Struct workflow
func FileYamlUnmarshal(data []byte) (*engine.Workflow, error) {

	// NOTE: modified by akshatm
	// fmt.Println(configFilePath)
	// filedata, err := ioutil.ReadFile(configFilePath)
	// if err != nil {
	// 	log.Fatalf("err1or: %v", err)
	// 	return err
	// }
	var t = engine.Workflow{}

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return &t, nil
}
