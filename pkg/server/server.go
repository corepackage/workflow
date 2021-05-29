package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/server/router"

	"github.com/gorilla/mux"
)

// Saving PId
func savePid(pid int) {
	file, err := os.Create(constants.PID_FILE)
	if err != nil {
		log.Printf("savePid : Unable to create a pid file : %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		log.Printf("savePid : Unable to write pid to the file : %v\n", err)
		os.Exit(1)
	}

	file.Sync() //flush to disk
}

// Start - to start the new mux server
func Start() error {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signalType := <-ch
		signal.Stop(ch)
		fmt.Println("Exit command received, Exiting...")

		fmt.Println("Received signal type : ", signalType)

		// Remove PID
		os.Remove(constants.PID_FILE)
		os.Exit(0)
	}()
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

func StartInBackground() {
	if _, err := os.Stat(constants.PID_FILE); err == nil {
		fmt.Println("Already running or /tmp/daemonize.pid file exists")
		os.Exit(1)
	}
	port := strconv.Itoa(models.EngConfig.Port)

	cmd := exec.Command(os.Args[0], constants.RUN_ENGINE, "--port="+port, "--path="+models.EngConfig.Prefix)
	cmd.Start()
	// fmt.Println("Daemon process ID is :", cmd.Process.Pid)
	savePid(cmd.Process.Pid)
	os.Exit(0)
}
