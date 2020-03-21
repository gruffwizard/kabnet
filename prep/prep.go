package prep


import (
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "strings"


)




// need to do vagrant box update

var serverRoot string = "https://mirror.openshift.com/pub/openshift-v4/"

var clientPath  string = "clients/ocp/"
var depPath    string = "dependencies/rhcos/latest/latest/"


//var client     string = "openshift-client-linux-4.3.3.tar.gz"
//var installer  string = "openshift-install-linux-4.3.3.tar.gz"

var initram string = "rhcos-4.3.0-x86_64-installer-initramfs.img"
var kernel  string = "rhcos-4.3.0-x86_64-installer-kernel"
var metal   string = "rhcos-4.3.0-x86_64-metal.raw.gz"


// /clients/ocp/latest/openshift-client-linux-4.3.1.tar.gz
// https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/openshift-install-linux-4.3.1.tar.gz

// https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/rhcos-4.3.0-x86_64-installer-initramfs.img
// https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/rhcos-4.3.0-x86_64-installer-kernel
// https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/rhcos-4.3.0-x86_64-metal.raw.gz

// assumes target dir already exists..

func Prepare(path string,cmdcfg defs.Config) {


  // create or update config  and do basic validation

  configNeedsUpdate,config:=buildConfig(path,cmdcfg)

  util.Section("OpenShift Cluster Configuration")
  util.Info("Domain Name         : %s ",config.Domain)
  util.Info("Cluster Name        : %s ",config.Cluster)
  util.Info("Openshift Version   : %s ",config.Version)
  util.Info("Openshift Installer : %s ",config.Tools)

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
