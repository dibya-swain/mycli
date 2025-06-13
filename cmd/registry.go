package cmd

import "github.com/spf13/cobra"

var RegisteredCommands []*cobra.Command

func Register(cmd *cobra.Command) {
	RegisteredCommands = append(RegisteredCommands, cmd)
}
