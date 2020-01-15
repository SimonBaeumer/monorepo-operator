package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	debug   bool
	version string
	ref     string
	build   string
)

const ConfigFile = ".monorepo-operator.yml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "monorepo-operator",
	Short: "Manage your monolithic repo and subtree splits",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// v contains the version number
// r the current commit ref
// b the current build number
func Execute(v string, b string, r string) {
	version = v
	ref = r
	build = b

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ConfigFile, "config file (default is .monorepo-operator.yml)")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in wd directory with name ".monorepo-operator" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName(ConfigFile)
	}

	viper.AutomaticEnv() // read in environment variables that match
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	}
}
