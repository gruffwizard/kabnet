package defs

import (

  "strings"
  "github.com/google/uuid"

)




type ManifestEntry struct {

    Remote string `yaml:"remote,omitempty"`
    Local  string `yaml:"local,omitempty"`

}

type Manifest struct {

  Entries []ManifestEntry `yaml:"entries,omitempty"`
}

type Options struct {
  UseVagrantDNS bool
}


type GeneratorDefinition struct {

  RootDir string `yaml:"root"`
  ImageDir string `yaml:"images"`
  Installer string `yaml:"installer"`
  SecretsDir string `yaml:"secrets"`
  ToolsDir string `yaml:"tools"`
  Installation string `yaml:"installation"`
  Pxe string `yaml:"pxe"`
  Cluster ClusterConfig `yaml:"cluster"`
  Secret string `yaml:"secret,omitempty"`
  Key    string `yaml:"key,omitempty"`
  Meta  *Config `yaml:"meta,omitempty"`
  Gateway string `yaml:"gateway"`
  Options  Options `yaml:"options,omitempty"`
  Mode  string `yaml:"mode,omitempty"`
}





func genMac() string {
    t:=  uuid.New().String()
    t=strings.Replace(t, "-", "",-1)
    return "00"+t[0:10]
}
