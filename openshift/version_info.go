package openshift

import (

  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/defs"
  "strings"
  "strconv"
  "regexp"
)

var baseOSURL     string = "https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/"
var baseToolsURL  string = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/"


func GetImageFileURLs(version string) (string,[]string) {

    pieces:=strings.Split(version, ".")
    majorVersion:=pieces[0]+"."+pieces[1]
    url:=baseOSURL+majorVersion+"/latest/"

    // get a list ..
    info:=readWebsiteLinks(url)
    // remove any that are not obviously not needed
    // so only those that end in "tar.gz" or  "-kernel" or ".img"

    var  candidates []string
    for _,v := range info {


        if  ( strings.HasSuffix(v,".img") || strings.HasSuffix(v,".raw.gz")  || strings.HasSuffix(v,"-kernel") ) {
             candidates=append(candidates,v)
        }
    }

    return url,candidates


}


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

func readWebsiteLinks(url string) []string {

  info:=util.ReadURLFile(url+"?F=0")
  reg := regexp.MustCompile(`href=".*"`)
  v:=reg.FindAllString(info,-1)


  var releases []string


  for _,a := range v {
      a=a[6:]
      a=a[0:len(a)-1]
      releases=append(releases,a)
  }

  return releases

}

func readWebsiteInfo(url string) []string {

  info:=util.ReadURLFile(url+"?F=0")
  reg := regexp.MustCompile(`href="([0-9].*)/"`)
  v:=reg.FindAllString(info,-1)


  var releases []string


  for _,a := range v {
      a=a[6:]
      a=a[0:len(a)-2]
      releases=append(releases,a)
  }

  return releases

}
/*
  Create an array of all the OpenShift V4 dependency sets  that exist online
  This is effectively the list of OpenShift versions available.
*/

func OpenShiftVersions() *defs.VersionInfo {

  info:=new(defs.VersionInfo)

  tools:=toolVersions()


  // top level
  majorVersions:=readWebsiteInfo(baseOSURL)

  var releases []*defs.Version

  for _,v := range majorVersions {
      version:=new(defs.Version)
      version.Name=v
      version.SubVersions=readWebsiteInfo(baseOSURL+v+"/")
      releases=append(releases,version)
  }

  info.Versions=releases
  info.Tools=releasedToolsOnly(tools)
  return info

}

func validateRelease(r string) bool {
  bits:=strings.Split(r,".")
  if len(bits)!=3 { return false}
  for _,b:= range bits {
    _,err := strconv.Atoi(b)
    if err!=nil  { return false }
  }
  return true
}

func releasedToolsOnly(tools []string) []string {

  var released []string

  for _,v:= range tools {
        if validateRelease(v) { released=append(released,v)}
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



func getOpenShiftFileList(path string) []string {

  releaseInfo:=util.ReadURLFile(path)
  bits:=strings.Split(releaseInfo,"\n")
  var files []string

  for _, v := range bits {
      parts:=strings.Split(v," ")
      if len(parts)>2 {  files=append(files,parts[2])}

  }

  return files

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
