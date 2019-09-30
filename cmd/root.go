package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "os"
)


var cfgFile string
var debug bool
const ConfigFile = ".monorepo-operator.yml"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "monorepo-operator",
  Short: "Manage your monolithic repo and subtree splits",
  Long: ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .monorepo-operator.yml)")
  rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "")

  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}
