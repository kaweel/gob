package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var info = &cobra.Command{
	Use:   "info",
	Short: "Show template info",
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

		infoFile := filepath.Join(root, cfg["info"])
		infoContent, err := os.ReadFile(infoFile)
		if err != nil {
			return err
		}
		fmt.Printf("%s : %s\n", name, string(infoContent))
		return nil
	},
}
