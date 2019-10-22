package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "The project command is used to operate on projects",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
