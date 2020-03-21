package util

import (
  "text/template"
    "bytes"
    "log"
)

func Parse(title string,c interface{},g string) string {

  t := template.New(title)
  _,err := t.Parse(g)
  if err!=nil {
    log.Fatal(err)
  }


  var buffer bytes.Buffer

  err2 := t.Execute(&buffer, c)

  if err2!=nil {
    log.Fatal(err2)
  }
  return buffer.String()

}

func Contains(c string,l []string) bool {
    for _,q := range l {
    //  log.Print(q,"--",c)
      if q==c {
        return true
      }
    }
    return false
}
