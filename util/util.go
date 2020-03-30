package util

import (
  "log"
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "encoding/json"
  "gopkg.in/yaml.v3"
  "strconv"
  "os/exec"


)

type Text struct {
  lines []string
}


func GetHomeDir() string {
  home,err := os.UserHomeDir()
  if err != nil { log.Fatal(err) }
  return home
}
func FindIDFile() (string,string) {

  home:=GetHomeDir()

  var ssh=home+"/.ssh"
  files, err := ioutil.ReadDir(ssh)
  if err != nil {
      log.Print("Can't find a local ID file")
      log.Fatal(err)
  }

  for _, file := range files {
      name:=file.Name()
      if strings.HasSuffix(name, ".pub") {
        // got a matching private key?
        priv := ssh+"/"+strings.TrimSuffix(name, ".pub")
          if _, err := os.Stat(priv); err == nil {
            return ssh+"/"+name , priv
          }
      }

  }
  return "",""
}
func WriteFile(dir string,fname string, contents string) {

  WriteFile2(dir+"/"+fname,contents)
}
func WriteFile2(path string, contents string) {

  log.Printf("save file %s",path)

 if path=="/tmp/foo/openshift/bootstrap.ign"  { log.Panic()}

  f, err := os.Create(path)
  if err!=nil { log.Fatal(err)}
  defer f.Close()
  f.WriteString(contents)
}

func ToString(n int) string {

  v := strconv.Itoa(n)

  return v

}


func ToInt(n string) int {

  v, err := strconv.Atoi(n)

  if err == nil {return v}

  return -1

}
func LoadAsYaml (path string, b interface{}) {

  yamlFile, err := os.Open(path)
  if  err!=nil  { log.Fatal(err)}
  defer yamlFile.Close()

  byteValue, _ := ioutil.ReadAll(yamlFile)
  err = yaml.Unmarshal([]byte(byteValue), b)
    if err != nil {
        log.Fatalf("cannot unmarshal data: %v", err)
    }

}

func SaveAsYaml (path string, d  interface{}) {
  data,err := yaml.Marshal(&d)
  if err != nil {
      log.Fatalf("cannot marshal data: %v", err)
  }
  WriteFile2(path,string(data))
}
func SaveAsJson (path string, d  interface{}) {
  data,err := json.Marshal(&d)
  if err != nil {
      log.Fatalf("cannot marshal data: %v", err)
  }
  WriteFile2(path,string(data))
}

func LoadFromJsonFile(ignfile string) map[string]interface{} {

  jsonFile, err := os.Open(ignfile)
  if  err!=nil  { log.Fatal(err)}
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    return  result
}

func CreateText() *Text {
  var t Text
  return &t
}

func LoadFile(file string) string {

  data, err := ioutil.ReadFile(file)

  if err!=nil {
    log.Printf("error loading file [%s]",file)
    log.Fatal(err)
  }
  return strings.TrimSuffix(string(data), "\n")

}

func (t *Text) AsString() string {
  return strings.Join(t.lines,"\n")+"\n"
}

func (t *Text) Add(format string, a ...interface{}) {
  t.lines=append(t.lines,fmt.Sprintf(format,a...))
}

func CopyFile(from string,todir string, tofile string) {

  log.Printf("copy %s to %s/%s",from,todir,tofile)

  f :=LoadFile(from)
  WriteFile(todir,tofile,f)

}

func MoveFile(from string, to string) {
  err := os.Rename(from,to)
	if err != nil {
		log.Fatal(err)
	}
}

func RecreateDir(path string,dirname string) string {

  dir := path+"/"+dirname

  err := os.RemoveAll(dir)

  if err!=nil {
    log.Fatal(err)
  }

  return CreateDir(path,dirname)

}

func CreateDir(path string,dirname string) string {
  dir := path+"/"+dirname

  if _, err := os.Stat(dir); os.IsNotExist(err) {
    os.Mkdir(dir, 0777)
  }
  return dir
}

func CreateDirFromPath(dir string) string {

  if _, err := os.Stat(dir); os.IsNotExist(err) {
    os.Mkdir(dir, 0777)
  }
  return dir
}

func FileMustExist(path string) {

  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("required file %s does not exist",path)
  }

}

func DirMustExist(path string) {

  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Fatalf("required directory %s does not exist",path)
  }

}

func FileExists(path string) bool {

  if _, err := os.Stat(path); os.IsNotExist(err) {
    return false
  }
  return true

}
func CreateFile(path string,fname string) *os.File {

  log.Printf("Create File %s/%s",path,fname)

  f, err := os.Create(path+"/"+fname)
  if err!=nil { log.Fatal(err)}
  return f

}

func Emit(f *os.File,format string, a ...interface{}) {

    _, err:=fmt.Fprintf(f,format+"\n",a...)
    if err != nil {
      log.Fatal(err)
    }
}

func FetchFiles(dir string, url string,files []string) {

    for k,n := range files {
      Info("File %d/%d",k+1,len(files))
      FetchFile(dir,url,n)
    }
}

func SpeculativeExecute(name string, args ...string) (string,error) {
  out, err := exec.Command(name,args...).Output()
  return strings.TrimSuffix(string(out), "\n"),err
}

func Execute(name string, args ...string) string {

  out, err := SpeculativeExecute(name,args...)
  if err != nil { log.Fatal(err)}
  return out
}
