package cmd

import (
	commoncmd "github.com/dibya-swain/commoncli/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "CLI that uses common and my cli",
}

func init() {

	_ = commoncmd.RegisteredCommands

	for _, cmd := range RegisteredCommands {
		RootCmd.AddCommand(cmd)
	}

	for _, cmd := range commoncmd.RegisteredCommands {
		RootCmd.AddCommand(cmd)
	}
}
