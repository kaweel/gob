package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var copy = &cobra.Command{
	Use:   "copy",
	Short: "Copy a project template and push to a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[TODO] Copy template and push to target repository")
	},
}
