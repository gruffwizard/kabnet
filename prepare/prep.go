package prep


import (
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "strings"


)

func Prepare(path string,cmdcfg defs.Config) {


  // create or update config  and do basic validation

  configNeedsUpdate,config:=buildConfig(path,cmdcfg)

  util.Section("OpenShift Cluster Configuration")
  util.Info("Domain Name         : %s ",config.Domain)
  util.Info("Cluster Name        : %s ",config.Cluster)
  util.Info("Openshift Version   : %s ",config.OpenshiftVersion)
  util.Info("Openshift Installer : %s ",config.ToolsVersion)

  checkForUserInstallableTools(path)

  checkForSecrets(path)

  checkForRequiredSupportFiles(path,config)

  checkForOpenShiftFiles(path,config)



    // if version is latest workout what that means



    buildManifest(config)


    util.Section("Check for cluster configuration file")

    configFile:=path+"/config.yaml"

    if util.FileExists(configFile)==false || configNeedsUpdate {
      util.Info("saving config")
      util.SaveAsYaml (configFile,config)
   }

    util.Section("Prep completed")
}

// vagrant box add --provider virtualbox generic/alpine310

func checkForLocalBox(name string) bool {

  list :=util.Execute("vagrant","box","list")

  return strings.Contains(list,name)

}

func addVagrantBox(name string) {

  util.Info("downloading vagrant box %s",name)
  util.Execute("vagrant","box","add","--provider","virtualbox",name)

}
