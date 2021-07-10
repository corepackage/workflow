package server

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/server/router"

	"github.com/gorilla/mux"
)

// Saving PId
func savePid(pid int) {
	pidFile := filepath.FromSlash(constants.PID_FILE)
	file, err := os.Create(pidFile)
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

	port := models.EngConfig.Port
	log.Println("starting engine on ", port)
	pidFile := filepath.FromSlash(constants.PID_FILE)

	r := mux.NewRouter()

	router.InitRoutes(r)
	// Converting port to string
	stringPort := strconv.Itoa(port)
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

	s := &http.Server{
		Addr:    ":" + stringPort,
		Handler: r,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

	}()

	signalType := <-ch
	fmt.Println("Exit command received, Exiting...")
	fmt.Println("Received signal type : ", signalType)

	// Gracefully shutting down server
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	// Remove PID
	os.Remove(pidFile)
	os.Exit(0)
	return nil
}

// StartInBackground : To start server in background
func StartInBackground() {
	pidFile := filepath.FromSlash(constants.PID_FILE)

	if _, err := os.Stat(pidFile); err == nil {
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

// Stop()
func Stop() error {
	pidFile := filepath.FromSlash(constants.PID_FILE)

	if _, err := os.Stat(pidFile); err == nil {
		data, err := ioutil.ReadFile(pidFile)
		if err != nil {
			fmt.Println("Not running")
			return errors.New("Server not running")
		}
		ProcessID, err := strconv.Atoi(string(data))

		if err != nil {
			fmt.Println("Unable to read and parse process id found in ", pidFile)
			return errors.New("Error stop server")
		}

		process, err := os.FindProcess(ProcessID)

		if err != nil {
			fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
			return errors.New("Error stop server")
		}
		// remove PID file
		os.Remove(pidFile)

		fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
		// kill process and exit immediately
		err = process.Kill()

		if err != nil {
			fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
			return errors.New("Error stop server")
		} else {
			fmt.Printf("Killed process ID [%v]\n", ProcessID)
			return nil
		}

	}

	fmt.Println("Not running.")
	return errors.New("Server not running")

}
