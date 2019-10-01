package cmd

import (
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "project exec [project-name] [command]",
	Short: "Executes a single command on a given project by project name",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		p, err := m.GetProject(args[0])
		if err != nil {
			log.Fatal(err)
		}

		m.ExecOnProject(p, strings.Join(args[1:], " "))
	},
}

func init() {
	projectCmd.AddCommand(execCmd)
}
