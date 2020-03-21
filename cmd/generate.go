package cmd

import (

  "github.com/spf13/cobra"
  "github.com/gruffwizard/kabnet/generate"
  "github.com/gruffwizard/kabnet/defs"
  "github.com/gruffwizard/kabnet/util"
  "path/filepath"
  "net"
  "log"

  "strings"
  "strconv"

)

var PConfigDir string = "."

// rootCmd represents the base command when called without any subcommands
var generateCmd = &cobra.Command{
  Use:   "gen",
  Short: "generate",
  Long: "Generate openshift cluster ",
  Run: func(cmd *cobra.Command, args []string) {

      if PConfigDir == "" { util.Fail("missing -d option")}
      config:= new(defs.Config)
      PConfigDir,_=filepath.Abs(PConfigDir)

      configFile:=PConfigDir+"/config.yaml"
      util.FileMustExist(configFile)
      util.LoadAsYaml(configFile,config)

      ip,netaddr,err := net.ParseCIDR(config.AddressPool)
      if err != nil { log.Fatal(err) }

      util.Section("configuration:")
      util.Info("Network: %s/%s",config.Cluster,config.Domain)
      util.Info("IP %s (%s)",ip,netaddr)
      util.Info("Nodes  : %d masters, %d workers",config.Masters,config.Workers)


      gen:= new(defs.GeneratorConfig)
      gen.Meta=config
      gen.Root=PConfigDir
      gen.Images=util.CreateDir(PConfigDir,"images")
      gen.Installer=util.CreateDir(PConfigDir,"installer")
      gen.Secrets=util.CreateDir(PConfigDir,"secrets")
      gen.Tools=util.CreateDir(PConfigDir,"tools")
      gen.Installation=util.CreateDir(PConfigDir,"installation")
      gen.Pxe=util.CreateDir(gen.Installation,"pxe")
      gen.Gateway="172.30.1.5"
      gen.Cluster.AddressPool=config.AddressPool
      gen.Cluster.Domain=config.Domain
      gen.Cluster.Cluster=config.Cluster

      // add physical machines..
      // if net mask is a zero set to 1
      ipliteral:=ip.To4().String()
      if strings.HasSuffix(ipliteral, ".0") { ipliteral=incIPAddress(ipliteral) }

      gen.Cluster.CreateKabnet(ipliteral)
      ipliteral=incIPAddress(ipliteral)

      gen.Cluster.CreateBootstrap(ipliteral)
      ipliteral=incIPAddress(ipliteral)

      for m := 0; m < config.Masters; m++ {
        gen.Cluster.AddMaster(ipliteral)
          ipliteral=incIPAddress(ipliteral)
      }
      for w := 0; w < config.Workers; w++ {
        gen.Cluster.AddWorker(ipliteral)
          ipliteral=incIPAddress(ipliteral)
      }

      // add aliases

      util.SaveAsYaml ("cluster.yaml",gen)

      generate.Generate(gen)

  },
}




func incIPAddress(ip string) string {
  split:=strings.Split(ip, ".")
  n,_:=strconv.Atoi(split[3])
  n++
  split[3]=strconv.Itoa(n)

  return strings.Join(split,".")

}
func init() {
        rootCmd.AddCommand(generateCmd)
        generateCmd.Flags().StringVarP(&PConfigDir, "dir", "d", "", "config directory (required)")

}
