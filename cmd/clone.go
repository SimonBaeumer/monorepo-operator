package cmd

import (
	"fmt"
	"github.com/SimonBaeumer/monorepo-operator/operator"
	"github.com/spf13/cobra"
	"log"
)

var reset bool

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clones all projects into the given operating directory. This command is used after a new clone of your git repository.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := operator.NewMonoRepo(cfgFile)
		if err != nil {
			log.Fatal(err)
		}

		if reset {
			fmt.Printf("> Removing operating directory at %s\n", m.OperatingDir)
			if err := m.Remove(); err != nil {
				log.Fatal(err)
			}
		}

		if err := m.Clone(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().BoolVarP(&reset, "reset", "r", false, "Remove the operating directory before cloning.")
}
