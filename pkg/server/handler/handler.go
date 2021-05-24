package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func WorkflowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Invoking " + vars["workflowId"]))
}

func ServerStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server running successfully"))
}
