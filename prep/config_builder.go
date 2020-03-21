package prep


import (

  "github.com/gruffwizard/kabnet/openshift"
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "github.com/asaskevich/govalidator"


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
      config.Version="latest"
      config.Tools  ="latest"
      config.Domain="kabnet.kab"
      config.Cluster="dev"
      config.Masters=3
      config.Workers=2
      config.AddressPool="192.168.50.0/24"
    }


    validateVersioninfo(&configNeedsUpdate,config,cmdcfg)

    validateNetworkConfig(&configNeedsUpdate,config,cmdcfg)


    return configNeedsUpdate,config
}

// RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
// ]{2,3})$`)

//         return RegExp.MatchString(domain)
// }


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

  if cmdcfg.Version!="" && cmdcfg.Version!=config.Version {
    config.Version=cmdcfg.Version
    *update=true
  }

  if cmdcfg.Tools!="" && config.Tools!=cmdcfg.Tools {
    config.Tools=cmdcfg.Tools
    *update=true
  }

  // available versions?
  util.Info("searching online version info for %s",config.Version)
  onlineVersions:=openshift.OpenShiftVersions()
  util.Info("available online versions are %v",onlineVersions)
  vinfo:=onlineVersions.GetVersion(config.Version)

  if vinfo==nil {
    util.Fail("unable to locate OpenShift Version %s",config.Version)
  }
  config.Version=vinfo.Name

  util.Info("Using OpenShift Version %s",config.Version)
  toolsVersion:=onlineVersions.GetTools(config.Tools)

  if toolsVersion=="" {
    util.Fail("unable to locate Tools Version %s",config.Tools)
  }
  config.Tools=toolsVersion

}
