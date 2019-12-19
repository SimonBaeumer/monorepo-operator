package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var operatingDir string
var clone bool

const defaultOperatingDir = ".git/.subtree-repos"

// initCmd represents the add command
var initCmd = &cobra.Command{
	Use:   "init [git base url] [path] [--dry-run]",
	Short: "Initializes repository by a given path",
	Long: `Init will take a path and register all directories under it as subtree-repos.
The git base url is the url where the subtree-splits will be synced.
As default the directory name will be used as the repository name.

Example usage:

	init git@github.com:SimonBaeumer path/to/repos
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("git base url and path are missing")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		monoRepo := operator.NewMonoRepoFromPath(args[0], args[1], operatingDir)
		if err := monoRepo.WriteConfigFile(cfgFile); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(operatingDir, 0755); err != nil {
			log.Fatal(err)
		}

		if clone {
			if err := monoRepo.Clone(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&operatingDir, "operating-dir", "o", defaultOperatingDir, "Set the directory with the original repos")
	initCmd.Flags().BoolVarP(&clone, "clone", "c", false, "Clones the repos directly")
}
