package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var structure = &cobra.Command{
	Use:   "struct",
	Short: "Show template structure",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := loadTemplateRoot()
		if err != nil {
			return err
		}
		name := args[0]
		cfgPath := filepath.Join(root, name+".json")
		b, err := os.ReadFile(cfgPath)
		if err != nil {
			return err
		}
		var cfg map[string]string
		json.Unmarshal(b, &cfg)

		structFile := filepath.Join(root, cfg["structure"])
		structContent, err := os.ReadFile(structFile)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(structContent))
		return nil
	},
}
