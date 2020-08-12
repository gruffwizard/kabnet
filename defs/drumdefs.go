package defs

import (
  "strings"
)
type KabnetConfig struct {
  config  []*ConfigElement  
}

func (d *KabnetConfig) Elements() []*ConfigElement {

  return d.config

}

func (d *KabnetConfig) PackageNames() []string {

      var names []string

      for _, v := range d.config {
          for _, p :=range v.package_list {
              names=append(names,p)
          }
      }
      return names
}

func (d *ConfigElement) Unit() string {
  return d.unit
}

func (d *ConfigElement) Commands() []string {

      return d.config_list
}


func (d *ConfigElement) Files() []*InstallFile {

      return d.files_list
}

func (d *ConfigElement) StartCommands() []string {

      return d.start_list
}


func (d *ConfigElement) AddCallback(f func(*ConfigElement,*GeneratorDefinition) )  {

  d.callback=f

}


func (d *KabnetConfig) Callback(g *GeneratorDefinition) {

      for _, v := range d.config {
          if v.callback!=nil    { v.callback(v,g) }
      }

}


func (d *KabnetConfig) Files() []*InstallFile{

      var files []*InstallFile

      for _, v := range d.config {
          for _, p :=range v.files_list {
              files=append(files,p)
          }
      }
      return files
}

var Kabnet *KabnetConfig = new(KabnetConfig)

type ConfigElement struct {
  unit string
  callback func(*ConfigElement,*GeneratorDefinition)
  package_list []string
  files_list   []*InstallFile
  config_list  []string
  start_list   []string
}

func (c *ConfigElement) PackageDesc() string  {

    if c.package_list==nil { return ""}

    return c.unit+" requires "+strings.Join(c.package_list," ")

}


func (c *ConfigElement) AddPackage(name string ) {
    c.package_list=append(c.package_list,name)
}

func (c *ConfigElement) AddStartCmd(name string ) {
  c.start_list=append(c.start_list,name)
}


func (c *ConfigElement) AddConfigCmd(name string ) {
  c.config_list=append(c.config_list,name)
}

func (c *ConfigElement) AddFile(local string,server string,content string ) {

  file:= new(InstallFile)
  file.LocalName=local
  file.Content=content
  file.ServerName =server

  c.files_list=append(c.files_list,file)

}

func NewConfigElement(unit string) *ConfigElement {
    ce :=new (ConfigElement)
    ce.unit=unit
    Kabnet.config=append(Kabnet.config,ce)
    return ce
}





type InstallFile struct {
    LocalName  string
    Content    string
    ServerName string
}
