package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var MainCommand = &cobra.Command{
	Use:   "matros",
	Short: "matros - workload orchestration",
}

func Execute() {
	MainCommand.Version = "0.1.0"

	// global flags
	MainCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose log output")

	// disable help command
	MainCommand.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	if err := MainCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
