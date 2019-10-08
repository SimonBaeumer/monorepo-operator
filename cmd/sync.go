package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

var useForce = false
// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Sync creates subtree splits and will sync them with the remote",
	Long: ``,
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

		m.Sync(args[0], useForce)
	},
}

func init() {
	syncCmd.Flags().BoolVarP(&useForce, "force", "f", false, "this will use the force flag in git commands")

	rootCmd.AddCommand(syncCmd)
}
