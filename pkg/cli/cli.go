package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/coredevelopment/workflow/internal/models"
	"github.com/coredevelopment/workflow/pkg/util"
)

// RunEngine - To start the workflow engine
func RunEngine() {

	// Getting flags respective to run command
	runSet := flag.NewFlagSet("", flag.ExitOnError)
	runSet.IntVar(&models.EngConfig.Port, "port", 7200, "Specific port to start the engine")
	runSet.IntVar(&models.EngConfig.Port, "p", 7200, "Specific port to start the engine")
	runSet.Parse(os.Args[2:])

	fmt.Println("starting engine on", models.EngConfig.Port)
}

// StopEngine - To stop the running instance of workflow engine
func StopEngine() {
	fmt.Println("Stopping the engine")
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
	fmt.Println(files)
}

// ListAll - to list all the existing workflow configurations
func ListAll() {}

func Remove() {}

func Purge() {}

func ShowInfo() {
	fmt.Println("Showing options")
}
