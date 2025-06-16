package cmd

import (
	"os"

	commoncmd "github.com/dibya-swain/commoncli/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "CLI that uses common and my cli",
}

func init() {
	// Default: commoncli commands override mycli on conflict
	// If COMMAND_OVERRIDE=mycli, then mycli commands override commoncli
	override := os.Getenv("COMMAND_OVERRIDE") // "mycli" or "commoncli"

	cmdMap := make(map[string]*cobra.Command)

	if override == "mycli" {
		// mycli overrides commoncli
		for _, cmd := range commoncmd.RegisteredCommands {
			cmdMap[cmd.Use] = cmd
		}
		for _, cmd := range RegisteredCommands {
			cmdMap[cmd.Use] = cmd
		}
	} else {
		// Default: commoncli overrides mycli
		for _, cmd := range RegisteredCommands {
			cmdMap[cmd.Use] = cmd
		}
		for _, cmd := range commoncmd.RegisteredCommands {
			cmdMap[cmd.Use] = cmd
		}
	}

	for _, cmd := range cmdMap {
		RootCmd.AddCommand(cmd)
	}
}
