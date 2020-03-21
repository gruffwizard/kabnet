package util

import (
  "text/template"
    "bytes"
    "log"
)



func WriteTemplate(g interface{},path string ,data string ) {

  engine  := template.New(path)

  var tpl bytes.Buffer
  t, terr := engine.Parse(data)

  if terr!= nil  {
      log.Print("error parsing "+path)
      log.Print(terr)
      log.Fatal("unable to create config")
  }
  if err := t.Execute(&tpl, g); err != nil {
      log.Print("error parsing "+path)
      log.Fatal(err)
  }

  WriteFile2(path,tpl.String())

}
