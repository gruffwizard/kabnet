
package openshift


import (
  "strings"
    "github.com/gruffwizard/kabnet/defs"
)

var baseToolsURL  string = "https://mirror.openshift.com/pub/openshift-v4/clients/ocp/"



func ClientVersions() []string {
  return readWebsiteInfo(baseToolsURL)
}



func ToClientLevel(level string) string {
    if(level=="")  { level="latest"}

    if(level=="latest") {
      versions:=ClientVersions()
      level=versions[len(versions)-1]
    }

    bits:=strings.Split(level,".")

    if len(bits)<3 { return "" }

    return level

}

func Clients(level string) *defs.OpenshiftTools {

    level=ToClientLevel(level)

    if level=="" {return nil}

    var tools defs.OpenshiftTools

    url:=baseToolsURL+level

    info:=readWebsiteLinks(url)

    for _,v := range info {

      if  strings.Contains(v,"client-linux") {tools.Client=v }
      if  strings.Contains(v,"install-linux"){ tools.Installer=v}

    }

    tools.BaseURL=url
    return &tools

}
