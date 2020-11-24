package cmd

import (
	"github.com/matrosio/matros/internal/initialize"
	"github.com/spf13/cobra"
)

var initnode bool
var initserver bool

var InitCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize matros",
	Long:  "Initialize matros",
	Run: func(cmd *cobra.Command, args []string) {
		// by default node and server initialization modes are false
		initnode = false
		initserver = false

		// type of initialize to execute
		// standalone - instance is both server and node
		// server 	  - instance is a server in a cluster
		// node   	  - instance is a node in a cluster
		initializetype := "standalone"

		// initialize server
		if initserver {
			initializetype = "server"
		}

		// initialize node
		if initnode {
			initializetype = "node"
		}

		initialize.Configure(initializetype)
	},
}

func init() {
	MainCommand.AddCommand(InitCommand)
	InitCommand.Flags().BoolVarP(&initserver, "server", "", false, "Initialize instance as server")
	InitCommand.Flags().BoolVarP(&initnode, "node", "", false, "Initialize instance as node")
}
