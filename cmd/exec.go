package cmd

import (
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"strings"
)

// operateCmd represents the operate command
var operateCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command on all your projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(ConfigFile)
		if err != nil {
			panic(err.Error())
		}

		m.Exec(strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(operateCmd)
}
