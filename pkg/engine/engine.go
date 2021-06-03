package engine

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/cryptography"
	"github.com/coredevelopment/workflow/pkg/db"
	"github.com/coredevelopment/workflow/pkg/parser"
)

// Init : to check status of workflow
func Init(workflowID string) (*models.Workflow, error) {
	config := db.GetActiveConfig(workflowID)

	// Checking workflow id
	if config.WorkflowID == "" {
		log.Println("Init : Workflow id is invalid")
		return nil, errors.New("Invalid workflow Id")
	}

	// Checking workflow status
	if !config.Active {
		log.Println("Init : Workflow is inactive")
		return nil, errors.New("Workflow is inactive")
	}

	// Decrypting configuration
	filename := config.WorkflowID + "_" + config.Version
	filePath := path.Join(filepath.FromSlash(constants.ENC_BASE_DIR), filename)
	byteData, err := cryptography.Decrypt(filePath)
	if err != nil {
		log.Println("Init : Error decrypting configuration", err)
		return nil, errors.New("Init : Error decrypting config")
	}
	var wf *models.Workflow
	// Parsing config
	if config.FileExt == ".yml" || config.FileExt == ".yaml" {
		wf, err = parser.FileYamlUnmarshal(byteData)
		if err != nil {
			log.Println("Init : Error parsing config")
			return nil, errors.New("Init : Error parsing config")
		}
	}
	return wf, nil
}

// Execute : Starting point of workflow engine
func Execute(wf *models.Workflow, headers map[string][]string, userData map[string]interface{}) (string, error) {
	fmt.Println(wf)
	fmt.Println(headers)
	fmt.Println(userData)
	return "", nil
}
