package openshift

import (

  "strings"
  "strconv"

)
//


func GetClientFileURLs(version string) (string,[]string) {

    url:=baseToolsURL+version

    candidates :=[]string {"openshift-client-linux.tar.gz","openshift-install-linux.tar.gz"}
    return url,candidates


}

/*
  Create an array of all the OpenShift V4 client tool versions that exist online
*/
func toolVersions() []string {

    return readWebsiteInfo(baseToolsURL)

}

/*
  Create an array of all the OpenShift V4 dependency sets  that exist online
  This is effectively the list of OpenShift versions available.
*/






func releasedToolsOnly(tools []string) []string {

  var released []string

  for _,v:= range tools {
        if validateRelease(v,3) { released=append(released,v)}
  }
  return released

}

// tools come from https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/sha256sum.txt
// deps (ie actual version ) comes from https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/sha256sum.txt
func OpenShiftToolsVersions1(version string) ([]string) {
  files:=getOpenShiftFileList("https://mirror.openshift.com/pub/openshift-v4/clients/ocp/"+version+"/sha256sum.txt")
  return files
}

func LatestOpenShiftVersion() (string) {

files:=getOpenShiftFileList("https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/sha256sum.txt")
return findVersion1(files)

}


// https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/latest/latest/sha256sum.txt

// guess version from file names

func findVersion1(files []string) (string) {



    for _, v := range files {
        bits:=strings.FieldsFunc(v, vsplit)
        version:=""
        for _,v := range bits {
          if _, err := strconv.Atoi(v); err == nil {
                if version!="" { version=version+"."}
                version=version+v
            }
        }

        if version!="" { return version}
    }

    return "<unknown>"


}




func vsplit(r rune) bool {
    return r == '-' || r == '.'
}
