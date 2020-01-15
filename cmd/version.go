package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", version)
		fmt.Printf("build: %s\n", build)
		fmt.Printf("ref: %s\n", ref)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
