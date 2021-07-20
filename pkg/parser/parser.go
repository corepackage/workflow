package parser

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/corepackage/workflow/internal/constants"
	"github.com/corepackage/workflow/pkg/cryptography"
	"github.com/corepackage/workflow/pkg/db"
	"github.com/corepackage/workflow/pkg/engine"
	"gopkg.in/yaml.v2"
)

var FilePath = RootDir()

var t = engine.Workflow{}

var (
	errInvalidID  = errors.New("GetWorkflowConfig : invalid workflow Id")
	errInactiveWF = errors.New("Workflow is inactive")
)

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

// 	workflowId, err := GetWorkflowId(configFilePath)
// 	fmt.Println(workflowId, err)
// 	workflowName, err := GetWorkflowName(configFilePath)
// 	fmt.Println(workflowName, err)

// }

// RootDir : To get the root directory of the workflow project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b), "../")
	return filepath.Dir(d)
}

// GetWorkflowConfig : to check status of workflow
func GetWorkflowConfig(workflowID string) (*engine.Workflow, error) {
	config := db.GetActiveConfig(workflowID)

	// Checking workflow id
	if config.WorkflowID == "" {
		log.Println("GetWorkflowConfig : Workflow id is invalid")
		return nil, errInvalidID
	}

	// Checking workflow status
	if !config.Active {
		log.Println("GetWorkflowConfig : Workflow is inactive")
		return nil, errInactiveWF
	}

	// Decrypting configuration
	filename := config.WorkflowID + "_" + config.Version
	filePath := path.Join(filepath.FromSlash(constants.ENC_BASE_DIR), filename)
	byteData, err := cryptography.Decrypt(filePath)
	if err != nil {
		log.Println("GetWorkflowConfig : Error decrypting configuration", err)
		return nil, errors.New("GetWorkflowConfig : Error decrypting config")
	}
	var wf *engine.Workflow
	// Parsing config
	if config.FileExt == ".yml" || config.FileExt == ".yaml" {
		wf, err = FileYamlUnmarshal(byteData)
		if err != nil {
			log.Println("GetWorkflowConfig : Error parsing config")
			return nil, errors.New("GetWorkflowConfig : Error parsing config")
		}
	}
	return wf, nil
}

// GetWorkflowId : To get the workflow Id from the file specified
func GetWorkflowId(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")

	// Getting decrypted data
	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err1or: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		_, err := FileYamlUnmarshal(filedata)
		if err != nil {
			log.Fatalf("Aborting : %v", err)
			return "", err
		}
	}
	if t.ID == "" {
		log.Fatalf("Aborting : %v", "Workflow Id is blank")
		return "", fmt.Errorf("Workflow Id is blank")
	}
	return t.ID, nil
}

// GetWorkflowVersion : To get the workflow version from the file specified
func GetWorkflowVersion(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err1or: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		_, err := FileYamlUnmarshal(filedata)
		if err != nil {
			log.Fatalf("Aborting : %v", err)
			return "", err
		}
	}
	if t.Version == "" {
		log.Printf("Aborting : %v", "Workflow version is blank")
		t.Version = "latest"
	}
	return t.Version, nil
}

// GetWorkflowName : To get the workflow Name from the file specified
func GetWorkflowName(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		_, err := FileYamlUnmarshal(filedata)
		if err != nil {
			log.Fatalf("Aborting : %v", err)
			return "", err
		}
	}
	if t.Name == "" {
		log.Fatalf("Aborting : %v", "Workflow Name is blank")
		return "", fmt.Errorf("Workflow Name is blank")
	}
	return t.Name, nil
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

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return &t, nil
}
