package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ByeCmd = &cobra.Command{
	Use:   "bye",
	Short: "bye from my CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bye from my CLI!")
	},
}

func init() {
	Register(ByeCmd)
}
