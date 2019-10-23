package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

// projectSplitCmd represents the projectSplit command
var projectSplitCmd = &cobra.Command{
	Use:   "split [project]",
	Short: "Create a subtree of a project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		p, err := m.GetProject(args[0])
		if err != nil {
			log.Fatal(err)
		}

		sha1 := m.SplitProject(p, "")
		fmt.Println(sha1)
	},
}

func init() {
	projectCmd.AddCommand(projectSplitCmd)
}
