package handler

import (
	"encoding/json"
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

	// Checking workflow status
	config, err := engine.Init(workflowID)
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

	// Marshalling user data to interface
	userData := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	headers := (map[string][]string)(r.Header)
	// Executing engine
	resp, err := engine.Execute(config, headers, userData)
	if err != nil {
		log.Println("Error executing workflow")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func ServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server running successfully"))
}
