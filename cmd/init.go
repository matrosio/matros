package cmd

import (
	"os"

	"github.com/matrosio/matros/internal/initialize"
	"github.com/spf13/cobra"
)

var initializetype string
var datadirectory string

var InitCommand = &cobra.Command{
	Use:   "init",
	Short: "initialize matros",
	Long:  "initialize matros",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize("standalone", &datadirectory)
		os.Exit(0)
	},
}

var InitServerCommand = &cobra.Command{
	Use:   "server",
	Short: "initialize server",
	Long:  "initialize server",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize("server", &datadirectory)
		os.Exit(0)
	},
}

var InitNodeCommand = &cobra.Command{
	Use:   "node",
	Short: "initialize node",
	Long:  "initialize node",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize("node")
		os.Exit(0)
	},
}

func init() {
	// set defaults
	datadirectory = "./matros"

	// flags
	InitCommand.Flags().StringVarP(&datadirectory, "data-dir", "", "", "directory to store matros data store")
	InitServerCommand.Flags().StringVarP(&datadirectory, "data-dir", "", "", "directory to store matros data store")
	InitCommand.MarkFlagRequired("data-dir")
	InitServerCommand.MarkFlagRequired("data-dir")

	// commands
	InitCommand.AddCommand(InitNodeCommand)
	InitCommand.AddCommand(InitServerCommand)
	MainCommand.AddCommand(InitCommand)
}
