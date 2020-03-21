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

    imageDir:=util.CreateDir(path,"images")
    util.Section("Check for openshift image files")
    url,names:=openshift.GetImageFileURLs(config.Version)

    util.FetchFiles(imageDir, url,names)

    util.Section("Check for openshift installer files")
    installerDir:=util.CreateDir(path,"installer")
    url,names=openshift.GetClientFileURLs(config.Tools)
    util.FetchFiles(installerDir, url,names)

    clientDir:=util.CreateDir(installerDir,"openshift-client")

    util.Section("extracting installers")
    err :=targz.Extract(installerDir+"/openshift-install-linux.tar.gz", installerDir)
    if err!=nil {  util.Fail("error extracting %s [%v]","openshift-install-linux.tar.gz",err)}

    targz.Extract(installerDir+"/openshift-client-linux.tar.gz", clientDir)

}
