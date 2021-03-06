package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/corepackage/workflow/pkg/engine"
	"github.com/corepackage/workflow/pkg/server/middleware"
	"github.com/corepackage/workflow/pkg/workflow"
	"github.com/gorilla/mux"
)

func WorkflowHandler(w http.ResponseWriter, r *http.Request) {
	reqCtx, ok := r.Context().Value(middleware.KeyRequest{}).(*http.Request)
	if !ok {
		log.Println("no context")
		http.Error(w, "no context", http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(reqCtx)
	// w.Write([]byte("Invoking " + vars["workflowId"]))
	workflowID := vars["workflowId"]
	if workflowID == "" {
		return
	}
	// Checking workflow status
	config, err := workflow.GetConfig(workflowID)
	if err != nil {
		log.Println("WorkflowHandler: Error initializing workflow")

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

	// Validating workflow access policy
	err = engine.Validate(reqCtx, w, config)
	if err != nil {
		log.Println("WorkflowHandler : validation failed")
		return
	}

	// getting query data

	queryParams := reqCtx.URL.Query()
	var bodyJson interface{}
	if r.Method != http.MethodGet {
		byteArray, err := ioutil.ReadAll(reqCtx.Body)
		if err != nil {
			log.Println("WorkflowHandler : Error fetching body json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request data"))
		}
		err = json.Unmarshal(byteArray, &bodyJson)
		if err != nil {
			log.Println("")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid request data"))
		}
	}
	resp, err := config.Run(r.Context(), reqCtx.Header, queryParams, bodyJson)
	if err != nil {
		log.Println("WorkflowHandler : Error running workflow")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}
	byteArray, err := json.Marshal(resp)
	if err != nil {
		log.Println("WorkflowHandler : error marshalling response ")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)

}

func ServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server running successfully"))
}
