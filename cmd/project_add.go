package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

var withoutClone = false

// projectAddCmd represents the projectAdd command
var projectAddCmd = &cobra.Command{
	Use:   "add [name] [git-url] [path]",
	Short: "Adds a project to the mono repo config",
	Long: `Add a project to the monorepo and clones it into the operating directory.

Example:
	project add new_project git@github.com:SimonBaeumer/monorepo-operator.git path/to/project
`,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		p := operator.Project{
			Name: args[0],
			GitUrl: args[1],
			Path: args[2],
		}

		m.Add(p)

		fmt.Printf("> Write config file %s\n", cfgFile)
		if err := m.WriteConfigFile(cfgFile); err != nil {
			log.Fatal(err)
		}

		if !withoutClone {
			fmt.Printf("> Cloning %s", p.Name)
			if err := p.GitClone(m.OperatingDir); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(projectAddCmd)
	projectCmd.Flags().BoolVarP(&withoutClone, "without-clone", "w", false, "Disables the cloning of the added project.")
}
