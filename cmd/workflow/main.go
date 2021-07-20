package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/corepackage/workflow/internal/constants"
	"github.com/corepackage/workflow/pkg/cli"
)

// Execute - To execute the specified option
func main() {

	// Retrieving operation from args
	if len(os.Args) <= 1 {
		cli.ShowHelp()
		return
	}
	command := os.Args[1]
	command = strings.ToLower(command)
	// Checking the execution command
	switch command {
	case constants.RUN_ENGINE:
		cli.RunEngine()
	case constants.STOP_ENGINE:
		cli.StopEngine()
	case constants.PUSH_CONFIG:
		cli.PushConfig()
	case constants.LIST_ALL_CONFIGS:
		cli.ListAll()
	case constants.REMOVE:
		cli.Remove()
	case constants.HELP:
		cli.ShowHelp()
	default:
		fmt.Println("Invalid command :", command)
	}
}
