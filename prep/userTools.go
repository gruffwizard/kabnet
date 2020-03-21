package prep

import (
  "github.com/gruffwizard/kabnet/util"
  //  "github.com/gruffwizard/kabnet/defs"

)

// check for those things the tool can't install


func checkForUserInstallableTools(path string) {

  passed:=true
  util.Section("Prereq tools check")

  info,err:=util.SpeculativeExecute("vagrant","-v")
  if err!=nil {
    util.Warn("missing vagrant");
    util.Warn("go to https://www.vagrantup.com/downloads.html to install")
    passed=false
  } else {
    util.Info("vagrant version : %s",info)

  }



  // install the 'Oracle VM VirtualBox Extension Pack'
  pullFile1:=path+"/pull-secret.txt"
  pullFile2:=path+"/secrets/pull-secret.txt"

  if util.FileExists(pullFile1)==false  && util.FileExists(pullFile2)==false {

        util.Warn("missing pull-secret.txt file.")
        util.Warn("Go to https://cloud.redhat.com/openshift/install/metal/user-provisioned and download a copy")
        passed=false
  } else {
      util.Info("pull secret file present")
  }

  if !passed {

    util.Fail("missing pre-req user installed tools required")
  }


}
