package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/coredevelopment/workflow/internal/models"
	"gopkg.in/yaml.v2"
)

var FilePath = RootDir()

var t = models.Workflow{}

func main() {

	configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	err := FileYamlUnmarshal(configFilePath)
	if err != nil {
		log.Fatalf("Aborting : %v", err)
	}

	fmt.Println("Primay Keys : ", t.PrimaryKey[0].PKey)
	fmt.Println("Workflow Name : ", t.Name)
	fmt.Println("Cors : ", t.CORS)
	// fmt.Println(t.Steps[1].ID)
	for key, value := range t.Steps {
		fmt.Println("Key :", key, "Value :", value)
		fmt.Println("Step Name:", value.Name)
		fmt.Println("Pre Condition:", value.PreCondition)
		fmt.Println("Pre Condition:", value.PostCondition)
	}

	workflowId, err := GetWorkflowId(configFilePath)
	fmt.Println(workflowId, err)
	workflowName, err := GetWorkflowName(configFilePath)
	fmt.Println(workflowName, err)

}

// RootDir : To get the root directory of the workflow project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../")
	return filepath.Dir(d)
}

// GetWorkflowId : To get the workflow Id from the file specified
func GetWorkflowId(filePath string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	err := FileYamlUnmarshal(filePath)
	if err != nil {
		log.Fatalf("Aborting : %v", err)
		return "", err
	}
	if t.ID == "" {
		log.Fatalf("Aborting : %v", "Workflow Id is blank")
		return "", fmt.Errorf("Workflow Id is blank")
	}
	return t.ID, nil
}

// GetWorkflowName : To get the workflow Name from the file specified
func GetWorkflowName(filePath string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	err := FileYamlUnmarshal(filePath)
	if err != nil {
		log.Fatalf("Aborting : %v", err)
		return "", err
	}
	if t.Name == "" {
		log.Fatalf("Aborting : %v", "Workflow Name is blank")
		return "", fmt.Errorf("Workflow Name is blank")
	}
	return t.Name, nil
}

// FileYamlUnmarshal : To unmarshal the YAML file to the Struct workflow
func FileYamlUnmarshal(configFilePath string) error {
	fmt.Println(configFilePath)
	filedata, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("err1or: %v", err)
		return err
	}

	err = yaml.Unmarshal(filedata, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}
	return nil
}
