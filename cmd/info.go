package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var infoName string

var info = &cobra.Command{
	Use:   "info",
	Short: "Show template info",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, err := loadTemplateRoot()
		if err != nil {
			return err
		}
		cfgPath := filepath.Join(root, infoName+".json")
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
		fmt.Printf("%s : %s\n", infoName, string(infoContent))
		return nil
	},
}

func init() {
	info.Flags().StringVar(&infoName, "name", "", "template name")
	info.MarkFlagRequired("name")
}
