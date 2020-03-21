package prep

import (
    "github.com/gruffwizard/kabnet/util"
    "github.com/gruffwizard/kabnet/defs"

)



var ipxeiso string = "ipxe.iso"
var ipxeisoServer string = "http://boot.ipxe.org/"


func checkForRequiredSupportFiles(path string,config *defs.Config) {

  util.Section("Check for required supporting files")

  imageDir:=util.CreateDir(path,"images")


    box:="generic/alpine310"
    if checkForLocalBox(box) {
      util.Info("vagrant box %s exists",box)
    } else {
      addVagrantBox(box)
    }

  


  util.FetchFile(imageDir,ipxeisoServer,ipxeiso)


}
