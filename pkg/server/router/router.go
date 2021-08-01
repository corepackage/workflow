package router

import (
	"path"

	"github.com/corepackage/workflow/pkg/engine"
	"github.com/corepackage/workflow/pkg/server/handler"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/", handler.ServerStatusHandler)
	// Adding route prefix
	pathPrefix := engine.GetEngConfig().Prefix
	if pathPrefix != "" {
		pathPrefix = "/" + pathPrefix
		pathPrefix = path.Clean(pathPrefix)
		dash := r.PathPrefix(pathPrefix).Subrouter()
		dash.HandleFunc("/{workflowId}", handler.WorkflowHandler)
	} else {
		r.HandleFunc("/{workflowId}", handler.WorkflowHandler)
	}

}
