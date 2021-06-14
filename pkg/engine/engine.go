package engine

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/cryptography"
	"github.com/coredevelopment/workflow/pkg/db"
	"github.com/coredevelopment/workflow/pkg/engine/auth"
	"github.com/coredevelopment/workflow/pkg/engine/cors"
	"github.com/coredevelopment/workflow/pkg/parser"
)

// GetWorkflowConfig : to check status of workflow
func GetWorkflowConfig(workflowID string) (*models.Workflow, error) {
	config := db.GetActiveConfig(workflowID)

	// Checking workflow id
	if config.WorkflowID == "" {
		log.Println("GetWorkflowConfig : Workflow id is invalid")
		return nil, errors.New("Invalid workflow Id")
	}

	// Checking workflow status
	if !config.Active {
		log.Println("GetWorkflowConfig : Workflow is inactive")
		return nil, errors.New("Workflow is inactive")
	}

	// Decrypting configuration
	filename := config.WorkflowID + "_" + config.Version
	filePath := path.Join(filepath.FromSlash(constants.ENC_BASE_DIR), filename)
	byteData, err := cryptography.Decrypt(filePath)
	if err != nil {
		log.Println("GetWorkflowConfig : Error decrypting configuration", err)
		return nil, errors.New("GetWorkflowConfig : Error decrypting config")
	}
	var wf *models.Workflow
	// Parsing config
	if config.FileExt == ".yml" || config.FileExt == ".yaml" {
		wf, err = parser.FileYamlUnmarshal(byteData)
		if err != nil {
			log.Println("GetWorkflowConfig : Error parsing config")
			return nil, errors.New("GetWorkflowConfig : Error parsing config")
		}
	}
	return wf, nil
}

// Init : Starting point of workflow engine
func Init(r *http.Request, w http.ResponseWriter, wf *models.Workflow) {
	// Marshalling user data to interface
	userData := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse user data"))
		return
	}
	// headers := (map[string][]string)(r.Header)
	err = cors.Validate(r, w, wf.CORS)
	if err != nil {
		log.Println("Error in CORS Policy")
		w.Write([]byte("Error in CORS Policy"))
		return
	}
	//TODO: Validate request if authorizer present
	err = auth.Validate(r, wf.Authorizer)
	if err != nil {
		log.Println("Request Not Valid")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Request Not Valid"))
		return
	}
	//TODO: Load the respective instance or create new instance
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Workflow executed successfully"))

}
