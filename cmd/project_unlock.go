package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

// unlockCmd represents the unlock command
var unlockCmd = &cobra.Command{
	Use:   "unlock [project]",
	Short: "Unlock unlocks a given repository from its current lock state",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		monoRepo, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		p, err := monoRepo.GetProject(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Unlock project " + p.Name)
		p.Unlock()
	},
}

func init() {
	projectCmd.AddCommand(unlockCmd)
}
