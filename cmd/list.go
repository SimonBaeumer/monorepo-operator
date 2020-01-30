package cmd

import (
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all projects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		header := []string{"name", "path", "git-url"}
		list := [][]string{}
		for _, p := range m.Projects {
			list = append(list, []string{p.Name, p.Path, p.GitUrl})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		table.AppendBulk(list)
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
