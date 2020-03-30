package prep

import (
    "github.com/gruffwizard/kabnet/util"
    "github.com/gruffwizard/kabnet/defs"
    "github.com/gruffwizard/kabnet/openshift"
      "github.com/walle/targz"
)


func checkForOpenShiftFiles(path string,config *defs.Config) {

    util.Section("Check for openshift files")


      // two types of files required  - install tools and runtime images

    imageDir:=util.CreateDir(path,"images/"+config.OpenshiftVersion)

    util.Section("Check for openshift image files")
    deps:=openshift.Dependencies(config.OpenshiftVersion)

    util.FetchFiles(imageDir, deps.BaseURL,deps.FileNames())

    util.Section("Check for openshift installer files")
    installerDir:=util.CreateDir(path,"installer/"+config.ToolsVersion)

    tools:=openshift.Clients(config.ToolsVersion)
    util.FetchFiles(installerDir, tools.BaseURL,tools.FileNames())

    clientDir:=util.CreateDir(installerDir,"openshift-client")

    util.Section("extracting installers")
    err :=targz.Extract(installerDir+"/"+tools.Installer, installerDir)
    if err!=nil {  util.Fail("error extracting %s [%v]",tools.Installer,err)}

    err=targz.Extract(installerDir+"/"+tools.Client, clientDir)
      if err!=nil {  util.Fail("error extracting %s [%v]",tools.Client,err)}

}
