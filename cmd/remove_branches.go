package cmd

import (
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

var (
	noLocal  bool
	noRemote bool
)

// removeBranchesCmd represents the removeBranches command
var removeBranchesCmd = &cobra.Command{
	Use:   "remove-branches",
	Short: "Removes branches in subtree repositories which do not exists in the monorepo",
	Long:  ``,
	Run: func(c *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(ConfigFile)
		if err != nil {
			log.Fatal(err)
		}

		m.RemoveBranches(noLocal, noRemote)
	},
}

func init() {
	rootCmd.AddCommand(removeBranchesCmd)
	removeBranchesCmd.Flags().BoolVarP(&noLocal, "no-local", "l", false, "Do not remove local branches")
	removeBranchesCmd.Flags().BoolVarP(&noRemote, "no-remote", "r", false, "Do not remove remote branches")
}
