package openshift

import (

  "github.com/gruffwizard/kabnet/util"
    "regexp"
      "strings"
        "strconv"

)
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



func validateRelease(r string,dots int) bool {
  bits:=strings.Split(r,".")
  if len(bits)!=dots { return false}
  for _,b:= range bits {
    _,err := strconv.Atoi(b)
    if err!=nil  { return false }
  }
  return true
}
