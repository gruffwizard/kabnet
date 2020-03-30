
package openshift

import (
  "strings"
    "github.com/gruffwizard/kabnet/defs"
)

var baseOSURL     string = "https://mirror.openshift.com/pub/openshift-v4/dependencies/rhcos/"



func ToLevel(level string) string {
    if(level=="")  { level="latest"}

    if(level=="latest") {
      versions:=OpenShiftMajorVersions()
      if len(versions)<1 { level=""
      } else {
      level=versions[len(versions)-1]
      }
    }

    bits:=strings.Split(level,".")

    if len(bits)<2 { return "" }

    if len(bits)==2 {
        subversions:=OpenShiftMinorVersion(level)
        if len(subversions)<1 { level=""
        } else {
          level=subversions[len(subversions)-1]
      }
    }

    return level

}

func Dependencies(level string) *defs.Dependencies {



    level=ToLevel(level)
    if level=="" { return nil }


  return getImageFileURLs(level)

}

func OpenShiftMajorVersions() []string {
  return readWebsiteInfo(baseOSURL)
}


func OpenShiftMinorVersion(major string) []string {


  if (!validateRelease(major,2)) { return []string{}}

  list:= readWebsiteInfo(baseOSURL+major+"/")

  if len(list)==0  { list=append(list,major+".0") }
  return list
}


func OpenShiftMinorVersionFiles(major string) []string {
  return readWebsiteInfo(baseOSURL+major+"/")
}




func getImageFileURLs(version string) (*defs.Dependencies) {

    var deps defs.Dependencies

    pieces:=strings.Split(version, ".")
    majorVersion:=pieces[0]+"."+pieces[1]
    url:=baseOSURL+majorVersion+"/"+version+"/"

    // get a list ..
    info:=readWebsiteLinks(url)

    if len(info)==0 {
      url=baseOSURL+majorVersion+"/latest/"
      info=readWebsiteLinks(url)
    }

    for _,v := range info {

      if  strings.HasSuffix(v,".img") {deps.Initramfs=v }
      if  strings.HasSuffix(v,".raw.gz")  { deps.Metal=v}
      if  strings.Contains(v,"-kernel")  { deps.Kernel=v}

    }

    deps.BaseURL=url
    return &deps


}
