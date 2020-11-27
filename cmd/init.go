package cmd

import (
	"os"

	"github.com/matrosio/matros/internal/initialize"
	"github.com/spf13/cobra"
)

var initializetype string

var InitCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize matros",
	Long:  "Initialize matros",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize("standalone")
		os.Exit(0)
	},
}

var InitServerCommand = &cobra.Command{
	Use:   "server",
	Short: "initialize server",
	Long:  "initialize server",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Initialize("server")
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
	InitCommand.AddCommand(InitNodeCommand)
	InitCommand.AddCommand(InitServerCommand)
	MainCommand.AddCommand(InitCommand)
}
