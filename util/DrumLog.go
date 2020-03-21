package util

import (
  "log"

)


func Section(data string,args ... interface{}) {

  log.Printf("\u001b[37m"+data+"\u001b[0m",args...)

}

func Info(data string,args ... interface{}) {
  log.Printf("\u001b[32m"+data+"\u001b[0m",args...)
}


func Warn(data string,args ... interface{}) {
  log.Printf("\u001b[31m"+data+"\u001b[0m",args...)
}


func Fail(data string,args ... interface{}) {
  log.Fatalf("\u001b[31m"+data+"\u001b[0m",args...)
}
