package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/server/router"

	"github.com/gorilla/mux"
)

// Start - to start the new mux server
func Start() error {
	port := models.EngConfig.Port
	log.Println("starting engine on ", port)

	r := mux.NewRouter()

	router.InitRoutes(r)
	// Converting port to string
	stringPort := strconv.Itoa(port)

	if err := http.ListenAndServe(":"+stringPort, r); err != nil {
		fmt.Println("Error starting workflow engine")
	}
	return nil
}
