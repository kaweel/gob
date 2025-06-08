package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var structureName string

var structure = &cobra.Command{
	Use:   "struct",
	Short: "Show template structure",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := loadTemplateRoot()
		if err != nil {
			return err
		}
		cfgPath := filepath.Join(root, structureName+".json")
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

func init() {
	structure.Flags().StringVar(&structureName, "name", "", "template name")
	structure.MarkFlagRequired("name")
}
