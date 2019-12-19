package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

var (
	useForce       = false
	removeBranches = false
	syncTag        = false
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Sync creates subtree splits and will sync them with the remote",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("a branch is needed to sync the subtree projects")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		if syncTag {
			m.SyncTag(args[0], useForce)
		} else {
			m.Sync(args[0], useForce)
		}

		// Remove branches if flag is set to true. Ignore local branches for syncing.
		if removeBranches {
			m.RemoveBranches(true, false)
		}
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&useForce, "force", "f", false, "this will use the force flag in git push")
	syncCmd.Flags().BoolVarP(&removeBranches, "remove-branches", "r", false, "removes branches in subtree-splits which do not exist on the mono-repo after syncing")
	syncCmd.Flags().BoolVarP(&syncTag, "tags", "t", false, "syncs a tag instead of a branch")

	rootCmd.AddCommand(syncCmd)
}
