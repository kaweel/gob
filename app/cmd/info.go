package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show information about a project template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[TODO] Show info about the template")
	},
}
