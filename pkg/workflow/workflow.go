package workflow

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"

	"github.com/corepackage/workflow/internal/constants"
	"github.com/corepackage/workflow/pkg/cryptography"
	"github.com/corepackage/workflow/pkg/db"
	"github.com/corepackage/workflow/pkg/engine"
	"github.com/corepackage/workflow/pkg/parser"
)

var (
	errInvalidID  = errors.New("GetConfig : invalid workflow Id")
	errInactiveWF = errors.New("Workflow is inactive")
)

// GetConfig : to check status of workflow
func GetConfig(workflowID string) (*engine.Workflow, error) {
	config := db.GetActiveConfig(workflowID)

	// Checking workflow id
	if config.WorkflowID == "" {
		log.Println("GetConfig : Workflow id is invalid")
		return nil, errInvalidID
	}

	// Checking workflow status
	if !config.Active {
		log.Println("GetConfig : Workflow is inactive")
		return nil, errInactiveWF
	}

	// Decrypting configuration
	filename := config.WorkflowID + "_" + config.Version
	filePath := path.Join(filepath.FromSlash(constants.ENC_BASE_DIR), filename)
	byteData, err := cryptography.Decrypt(filePath)
	if err != nil {
		log.Println("GetConfig : Error decrypting configuration", err)
		return nil, errors.New("GetConfig : Error decrypting config")
	}
	var wf *engine.Workflow
	// Parsing config
	if config.FileExt == ".yml" || config.FileExt == ".yaml" {
		wf, err = parser.FileYamlUnmarshal(byteData)
		if err != nil {
			log.Println("GetConfig : Error parsing config")
			return nil, errors.New("GetConfig : Error parsing config")
		}
	}
	return wf, nil
}

// GetID : To get the workflow Id from the file specified
func GetID(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	var t *engine.Workflow
	// Getting decrypted data
	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err1or: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		t, err = parser.FileYamlUnmarshal(filedata)
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

// GetVersion : To get the workflow version from the file specified
func GetVersion(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	var t *engine.Workflow

	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("err1or: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		t, err = parser.FileYamlUnmarshal(filedata)
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

// GetName : To get the workflow Name from the file specified
func GetName(filePath, ext string) (string, error) {
	// configFilePath := path.Join(FilePath, "configs/workflow_config.yml")
	var t *engine.Workflow

	filedata, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error: %v", err)
		return "", err
	}
	// Checking file extension
	if ext == ".yaml" || ext == ".yml" {
		t, err = parser.FileYamlUnmarshal(filedata)
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
