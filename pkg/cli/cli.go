package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/cryptography"
	"github.com/coredevelopment/workflow/pkg/db"
	"github.com/coredevelopment/workflow/pkg/parser"
	"github.com/coredevelopment/workflow/pkg/server"
	"github.com/coredevelopment/workflow/pkg/util"
)

// RunEngine - To start the workflow engine
func RunEngine() {

	// Getting flags respective to run command
	runSet := flag.NewFlagSet("", flag.ExitOnError)
	runSet.IntVar(&models.EngConfig.Port, constants.PORT, 7200, "Specific port to start the engine")
	runSet.IntVar(&models.EngConfig.Port, constants.PORT_SHORT, 7200, "Specific port to start the engine")
	runSet.StringVar(&models.EngConfig.Prefix, constants.PATH, "", "Specific port to start the engine")

	// Check to see if run in detach mode
	var isDetached bool
	runSet.BoolVar(&isDetached, constants.DETACH, false, "Specific port to start the engine")
	runSet.BoolVar(&isDetached, constants.DETACH_SHORT, false, "Specific port to start the engine")
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

	// Stopping all configs and engine
	if len(os.Args) <= 2 {
		fmt.Println("Stopping the engine")
		if err := server.Stop(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	} else {
		// Stopping specific config
		workflows := os.Args[2:]
		for _, w := range workflows {
			arr := strings.Split(w, ":")
			var id, version string
			if len(arr) <= 1 {
				id = arr[0]
				version = "latest"
			} else {
				id, version = arr[0], arr[1]
			}
			fmt.Printf("Stopping %v\n", id+":"+version)
			err := db.DeactivateConfig(id, version)
			if err != nil {
				fmt.Println("Error stopping config", err)
			}
			fmt.Println("Configuration stopped")
		}
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
		ext := filepath.Ext(file)
		// Get the respective data to store in db
		workflowId, err := parser.GetWorkflowId(file, ext)
		if err != nil {
			log.Printf("Workflow Id not found for file %s: %v\n", file, err)
			continue
		}
		workflowName, err := parser.GetWorkflowName(file, ext)
		if err != nil {
			log.Printf("Workflow Name not found for file %s: %v\n", file, err)
			continue
		}
		version, err := parser.GetWorkflowVersion(file, ext)
		if err != nil {
			log.Printf("Workflow Version not found taking \"latest\"\n")
		}

		// Encrypting data
		cryptErr := cryptography.Encrypt(file, path.Join(filepath.FromSlash(constants.ENC_BASE_DIR), workflowId+"_"+version))
		if cryptErr != nil {
			log.Printf("Error while encrypting file %v : %v", file, cryptErr)
			continue
		}

		// Updating database
		if err := db.InsertOrUpdateConfig(workflowId, workflowName, version, ext); err != nil {
			os.Exit(1)
		}
	}
}

// ListAll - to list all the existing workflow configurations
func ListAll() {
	// checking args
	var showActive bool
	runSet := flag.NewFlagSet("", flag.ExitOnError)
	runSet.BoolVar(&showActive, constants.ALL, false, "Active workflow flag")
	runSet.BoolVar(&showActive, constants.ALL_SHORT, false, "Active workflow flag")
	runSet.Parse(os.Args[2:])

	// Showing only active configs
	var configs []models.WorkflowConfig
	if !showActive {
		configs = db.GetActiveConfigs()
	} else {
		configs = db.GetAllConfig()
	}

	// Printing the configs
	fmt.Printf("Workflow Id\t\tWorkflow Name\t\tVersion\t\tCreated At\t\tUpdated At\t\t\n")
	fmt.Printf("------------------------------------------------------------------------------------------------\n")
	for _, c := range configs {
		fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\t\t\n", c.WorkflowID, c.WorkflowName, c.Version, c.CreatedAt, c.UpdatedAt)
	}
}
func Remove() {
	if len(os.Args) <= 2 {
		RemoveAll()
	} else {
		// Stopping specific config
		workflows := os.Args[2:]
		for _, w := range workflows {
			arr := strings.Split(w, ":")
			var id, version string
			if len(arr) <= 1 {
				id = arr[0]
				version = "latest"
			} else {
				id, version = arr[0], arr[1]
			}
			fmt.Printf("Deleting %v\n", id+":"+version)
			err := db.DeleteConfig(id, version)
			if err != nil {
				fmt.Println("Error deleted config", err)
			}
			fmt.Println("Configuration deleted")
		}
	}
}

func RemoveAll() {}

func ShowHelp() {
	fmt.Println("Showing options")
}
