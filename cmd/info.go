package cmd

import (

  "github.com/spf13/cobra"
  "github.com/gruffwizard/kabnet/openshift"
  "fmt"

)
//FindAllStringSubmatch
var infoCmd = &cobra.Command{
  Use:   "info",
  Short: "info",
  Long: "List information about available OpenShift versions",
}

var infoLatestCmd = &cobra.Command{
  Use:   "latest",
  Short: "latest",
  Long: "List information about latest OpenShift versions",

  Run: func(cmd *cobra.Command, args []string) {

      versions:=openshift.OpenShiftVersions()
      latest:=versions.GetVersion("latest")
      tlatest:=versions.GetTools("latest")
      fmt.Printf("\nLatest Online Version with tools : [%s] , [%s]\n",latest,tlatest)


  },
}

var infoAllCmd = &cobra.Command{
  Use:   "all",
  Short: "all",
  Long: "List all information about available OpenShift versions",

  Run: func(cmd *cobra.Command, args []string) {

    versions:=openshift.OpenShiftVersions()

    fmt.Printf("\nOnline Versions:")

    for _, v := range versions.Versions {
      fmt.Printf("\n  [%s]",v.Name)
      fmt.Printf("\n      Subversions : ")
      for _,t := range v.SubVersions {
        fmt.Printf("[%s] ",t)
      }

    }
    fmt.Printf("\n")
    fmt.Printf("\n      Tools       : ")
    for _,t := range versions.Tools {
      fmt.Printf("[%s] ",t)
    }
  },
}

func init() {
        rootCmd.AddCommand(infoCmd)
        infoCmd.AddCommand(infoAllCmd)
        infoCmd.AddCommand(infoLatestCmd)

}
