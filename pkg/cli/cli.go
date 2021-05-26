package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/cryptography"
	fileops "github.com/coredevelopment/workflow/pkg/fileOps"
	"github.com/coredevelopment/workflow/pkg/server"
	"github.com/coredevelopment/workflow/pkg/util"
)

// RunEngine - To start the workflow engine
func RunEngine() {

	// Getting flags respective to run command
	runSet := flag.NewFlagSet("", flag.ExitOnError)
	runSet.IntVar(&models.EngConfig.Port, "port", 7200, "Specific port to start the engine")
	runSet.IntVar(&models.EngConfig.Port, "p", 7200, "Specific port to start the engine")
	runSet.StringVar(&models.EngConfig.Prefix, "path", "", "Specific port to start the engine")

	// Check to see if run in detach mode
	var isDetached bool
	runSet.BoolVar(&isDetached, "detach", false, "Specific port to start the engine")
	runSet.BoolVar(&isDetached, "d", false, "Specific port to start the engine")
	runSet.Parse(os.Args[2:])

	// TODO: Handler running in background
	// Starting instances of workflow server
	var err error
	if isDetached {
		server.StartInBackground()
	} else {
		err = server.Start()
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}

// StopEngine - To stop the running instance of workflow engine
func StopEngine() {

	// TODO: Added user logs
	fmt.Println("Stopping the engine")
	if _, err := os.Stat(constants.PIDFile); err == nil {
		data, err := ioutil.ReadFile(constants.PIDFile)
		if err != nil {
			fmt.Println("Not running")
			os.Exit(1)
		}
		ProcessID, err := strconv.Atoi(string(data))

		if err != nil {
			fmt.Println("Unable to read and parse process id found in ", constants.PIDFile)
			os.Exit(1)
		}

		process, err := os.FindProcess(ProcessID)

		if err != nil {
			fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
			os.Exit(1)
		}
		// remove PID file
		os.Remove(constants.PIDFile)

		fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
		// kill process and exit immediately
		err = process.Kill()

		if err != nil {
			fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
			os.Exit(1)
		} else {
			fmt.Printf("Killed process ID [%v]\n", ProcessID)
			os.Exit(0)
		}

	} else {

		fmt.Println("Not running.")
		os.Exit(1)
	}
}

// PushConfig - To push the workflow config to the engine
func PushConfig() {
	files := make([]string, 0)
	if len(os.Args) <= 2 {
		fmt.Println("Please provide a file")
		return
	} else {

		for _, file := range os.Args[2:] {
			err := util.ValidateFile(file)
			if err != nil {
				fmt.Println(err)
			} else {
				files = append(files, file)
			}
		}
	}

	// Ranging over files and invoking encryption
	for _, file := range files {
		log.Println("Encrypting File: ", file)
		cryptErr := cryptography.Encrypt(file, path.Join(constants.ENC_BASE_DIR, fileops.GetFileName(file)))
		if cryptErr != nil {
			log.Printf("Error while encrypting file %v : %v", file, cryptErr)
		}
	}
}

// ListAll - to list all the existing workflow configurations
func ListAll() {}

func Remove() {}

func RemoveAll() {}

func ShowHelp() {
	fmt.Println("Showing options")
}
