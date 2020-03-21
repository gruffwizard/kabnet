
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
// kabnet fetch pulls down all the files needed
// kabnet generate creates all the config files needed .
// kabnet build drives the code to create the cluster..


var rootCmd = &cobra.Command{
  Use:   "kabnet",
  Short: "OpenShift cluster generator ",
  Long: `Generates deployment config and related scripts to
         deploy openshift cluster`,

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
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {

}
