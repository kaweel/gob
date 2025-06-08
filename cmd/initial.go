package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initial = &cobra.Command{
	Use:   "init",
	Short: "set working directory for gob templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		var abs string
		var base string

		defer func() {
			fmt.Printf("gob root path: %s\n", base)
		}()

		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}

		abs = filepath.Join(home, ".gob", "setting.json")
		base = filepath.Dir(abs)
		if _, err := os.Stat(base); os.IsNotExist(err) {
			fmt.Printf("creating root director %v\n", base)
			if err := os.MkdirAll(base, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}

		templates := fmt.Sprintf("%s/templates", base)
		if _, err := os.Stat(templates); os.IsNotExist(err) {
			fmt.Printf("creating templates director %v\n", templates)
			if err := os.MkdirAll(templates, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		}

		if _, err := os.Stat(abs); os.IsNotExist(err) {
			data := map[string]string{
				"templates": templates,
				"version":   "v0.0.0",
			}
			formatted, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal: %w", err)
			}

			if err := os.WriteFile(abs, formatted, 0644); err != nil {
				return fmt.Errorf("failed to create setting.json: %w", err)
			}
		}

		return nil
	},
}

func loadTemplateRoot() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, ".gob", "setting.json")

	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("missing .gob/setting.json, please run 'gob root' :%w", err)
	}
	var cfg map[string]string
	json.Unmarshal(b, &cfg)
	return cfg["templates"], nil
}
