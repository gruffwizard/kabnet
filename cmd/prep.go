package cmd

import (

  "github.com/spf13/cobra"
  "github.com/gruffwizard/kabnet/prep"
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "path/filepath"
  "log"
)

var cmdLineParams defs.Config
var PDirectory string


// rootCmd represents the base command when called without any subcommands
var prepCmd = &cobra.Command{
  Use:   "prep",
  Short: "prep",
  Long: "prep required files and images ",
  Run: func(cmd *cobra.Command, args []string) {

    if PDirectory == "" { util.Fail("missing -d option")}



    path,err:=filepath.Abs(PDirectory)
    if err!=nil {log.Panic(err)}
    if util.FileExists(path) ==false { util.CreateDirFromPath(path) }

    prep.Prepare(path,cmdLineParams)


  },
}



func init() {
        rootCmd.AddCommand(prepCmd)
        prepCmd.Flags().StringVarP(&PDirectory,              "dir", "d", "", "output directory (required)")
        prepCmd.Flags().StringVarP(&cmdLineParams.Version,       "version",  "v", "", "OpenShift version (defaults to latest)")
        prepCmd.Flags().StringVarP(&cmdLineParams.Tools,  "tools",    "t", "", "OpenShift tools version (defaults to latest within version)")
        prepCmd.Flags().StringVarP(&cmdLineParams.Domain,     "domain",   "m", "", "OpenShift Domain (defaults to kabnet.kab)")
          prepCmd.Flags().StringVarP(&cmdLineParams.Cluster,     "cluster",   "c", "", "OpenShift Cluster (defaults to dev)")
}
