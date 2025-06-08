package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var commands = &cobra.Command{
	Use:   "gob",
	Short: "gob is a CLI tool to manage project templates",
	Long:  `gob helps you list, inspect, clone, and push project templates using a simple CLI interface.`,
}

func init() {
	commands.AddCommand(initial)
	commands.AddCommand(list)
	commands.AddCommand(info)
	commands.AddCommand(structure)
	commands.AddCommand(copy)
}

func Execute() {
	if err := commands.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
