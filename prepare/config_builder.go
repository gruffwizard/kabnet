package prep


import (

  "github.com/gruffwizard/kabnet/openshift"
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "github.com/asaskevich/govalidator"
  "fmt"


)

// Builds a valid config (or reports errors)
// returns true if the ondisk version needs updating / creating
func buildConfig(path string,cmdcfg defs.Config) (bool,*defs.Config) {

    util.Section("Validating config")

    configNeedsUpdate:=false

    config:= new(defs.Config)
    // do we have an existing config file
    configFile:=path+"/config.yaml"

    if util.FileExists(configFile) {
      util.LoadAsYaml(configFile,config)
    } else {
      config.OpenshiftVersion="4.3.8"
      config.ToolsVersion  ="4.3.8"
      config.Domain="kabnet.kab"
      config.Cluster="dev"
      config.Masters=3
      config.Workers=2
      config.AddressPool="192.168.50.0/24"
    }


    validateVersioninfo(&configNeedsUpdate,config,cmdcfg)

    validateNetworkConfig(&configNeedsUpdate,config,cmdcfg)

    validateDependencies(&configNeedsUpdate,config)
    validateTools(&configNeedsUpdate,config)

    return configNeedsUpdate,config
}

func validateDependencies(update *bool,config *defs.Config) {

  files:=openshift.Dependencies(config.OpenshiftVersion)

  compareAndSet(update,&config.ImageInfo.BaseURL,&files.BaseURL)
  compareAndSet(update,&config.ImageInfo.Initramfs,&files.Initramfs)
  compareAndSet(update,&config.ImageInfo.Kernel,&files.Kernel)
  compareAndSet(update,&config.ImageInfo.Metal ,&files.Metal )

}


func validateTools(update *bool,config *defs.Config) {


  files:=openshift.Clients(config.ToolsVersion)

  compareAndSet(update,&config.OpenshiftTools.BaseURL,&files.BaseURL)
  compareAndSet(update,&config.OpenshiftTools.Client,&files.Client)
  compareAndSet(update,&config.OpenshiftTools.Installer,& files.Installer)


}



func validateNetworkConfig(update *bool,config *defs.Config,cmdcfg defs.Config) {

    if cmdcfg.Domain!="" && config.Domain!=cmdcfg.Domain {
      config.Domain=cmdcfg.Domain
      *update=true
    }

    if cmdcfg.Cluster!="" && config.Cluster!=cmdcfg.Cluster {
      config.Cluster=cmdcfg.Cluster
      *update=true
    }

    if !govalidator.IsDNSName(config.Domain) {
      util.Fail("Domain name [%s] is invalid",config.Domain)
    }

    if !govalidator.IsDNSName(config.Cluster+"."+config.Domain) {
      util.Fail("Cluster name [%s] is invalid",config.Cluster)
    }


}

func validateVersioninfo(update *bool ,config *defs.Config,cmdcfg defs.Config ) {

  // if we've been given command line versions use them

  if cmdcfg.OpenshiftVersion!="" && cmdcfg.OpenshiftVersion!=config.OpenshiftVersion {
    config.OpenshiftVersion=cmdcfg.OpenshiftVersion
    *update=true
  }

  if cmdcfg.ToolsVersion!="" && config.ToolsVersion!=cmdcfg.ToolsVersion {
    config.ToolsVersion=cmdcfg.ToolsVersion
    *update=true
  }

  // available versions?
  util.Info("searching online version info for %s",config.OpenshiftVersion)
  vinfo:=openshift.ToLevel(config.OpenshiftVersion)

  if vinfo=="" {
    util.Fail("unable to locate OpenShift Version %s",config.OpenshiftVersion)
  }

  config.OpenshiftVersion=vinfo
  util.Info("Using OpenShift Version %s",config.OpenshiftVersion)

  toolsVersion:=openshift.ToClientLevel(config.ToolsVersion)

  if toolsVersion=="" {
    util.Fail("unable to locate Tools Version %s",config.ToolsVersion)
  }
  config.ToolsVersion=toolsVersion

}

func compareAndSet(f *bool,original *string, input *string) {

    fmt.Printf("%s==%s",*original,*input)

    if *original!=*input {
        *original=*input
        *f=true
    }
}
