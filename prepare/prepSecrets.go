
package prep

import (
  "github.com/gruffwizard/kabnet/util"
  
)



func checkForSecrets(path string) {

// file checks.
util.Section("Check for required secrets")

// got a local pull-secret.txt file?
secDir:=util.CreateDir(path,"secrets")
pullFile:=secDir+"/pull-secret.txt"

if util.FileExists(pullFile)==false {
    topCopy:=path+"/pull-secret.txt"
    if util.FileExists(topCopy) {
      util.MoveFile(topCopy,pullFile)
    } else {
      util.Fail("missing pull-secret.txt file.\n Go to https://cloud.redhat.com/openshift/install/metal/user-provisioned and download a copy")
    }
}
util.Info("pull secret file exists")
  sshFile:=secDir+"/core_rsa"
if util.FileExists(sshFile)==false {
  util.Info("creating ssh key")
  util.Execute("ssh-keygen","-t","rsa","-b","4096","-N","kabnet","-f",sshFile)
} else {
  util.Info("ssh key exists")
}

}
