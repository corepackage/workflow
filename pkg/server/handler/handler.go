package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/coredevelopment/workflow/pkg/engine"
	"github.com/gorilla/mux"
)

func WorkflowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// w.Write([]byte("Invoking " + vars["workflowId"]))
	workflowID := vars["workflowId"]
	if workflowID == "" {
		return
	}
	// Checking workflow status
	config, err := engine.GetWorkflowConfig(workflowID)
	if err != nil {
		log.Println("Error initializing workflow")

		errString := err.Error()
		if strings.Contains(errString, "Init") {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error executing workflow"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errString))
		}
		return
	}

	// Executing engine
	engine.Init(r, w, config)

}

func ServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server running successfully"))
}
