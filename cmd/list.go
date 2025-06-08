package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := loadTemplateRoot()
		if err != nil {
			return err
		}
		entries, err := os.ReadDir(root)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			fmt.Printf("no template available\n")
			return nil
		}
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
				name := strings.TrimSuffix(entry.Name(), ".json")
				fmt.Printf("%s\n", name)
			}
		}
		return nil
	},
}
