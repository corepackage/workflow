package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/coredevelopment/workflow/internal/constants"
	"github.com/coredevelopment/workflow/pkg/cli"
)

// Execute - To execute the specified option
func Execute() {

	// Retrieving operation from args
	if len(os.Args) <= 1 {
		cli.ShowInfo()
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
	case constants.HELP:
		cli.ShowInfo()
	default:
		fmt.Println("Invalid command :", command)
	}
}
