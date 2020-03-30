package defs

import (
  "strconv"
  "strings"
  "github.com/google/uuid"

)
var sshPortRef int = 1200


type Dependencies struct {
    BaseURL string
    Initramfs string
    Kernel string
    Metal  string

}


func (d *Dependencies) FileNames() []string {

  return []string{ d.Initramfs,d.Kernel,d.Metal}
}

type OpenshiftTools struct {
    BaseURL string
    Client string
    Installer string

}

func (d *OpenshiftTools) FileNames() []string {

  return []string{ d.Client,d.Installer}
}


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

type Config struct {
  OpenshiftVersion string `yaml:"openshift"`
  ToolsVersion string   `yaml:"tools"`
  Domain  string `yaml:"domain"`
  Cluster string `yaml:"cluster"`

  AddressPool string `yaml:"addresspool"`
  Masters int    `yaml:"masters"`
  Workers int    `yaml:"workers"`

  ImageInfo Dependencies `yaml:"deps,omitempty"`
  OpenshiftTools OpenshiftTools `yaml:"ostools,omitempty"`

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

}

type Machine struct {
    Name string `yaml:"name"`
    IP  string   `yaml:"ip"`
    Mac string `yaml:"mac"`
    MacVB string `yaml:"macvb,omitempty"`
    SSHProxyPort int `yaml:"sshproxyport"`

}
type ClusterConfig struct {
    Domain  string `yaml:"domain"`
    Cluster string `yaml:"cluster"`
    SubnetIP string `yaml:"subnetip,omitempty"`
    AddressPool string `yaml:"address_pool"`
    AddressRange []string `yaml:"address_range,omitempty"`
    NetMask string `yaml:"netmask,omitempty"`
    Kabnet *Machine  `yaml:"kabnet"`
    BootStrap *Machine `yaml:"bootstrap"`
    Masters  []*Machine `yaml:"masters"`
    Workers  []*Machine `yaml:"workers"`


}


func (c *ClusterConfig) AddMaster(ip string) *Machine {

    m :=new(Machine)

    mid :=len(c.Masters)+1
    m.Name="master"+strconv.Itoa(mid)
    m.IP=ip
    sshPortRef++
    m.SSHProxyPort=sshPortRef
    m.MacVB=genMac()
    m.Mac=toForm(m.MacVB,":")
    c.Masters=append(c.Masters,m)
    return m

}


func toForm(mac string,s string) string {

    r := mac[0:2] + s + mac[2:4] + s + mac[4:6] + s + mac[6:8] + s + mac[8:10] + s + mac[10:12]
    return r

}



func genMac() string {
    t:=  uuid.New().String()
    t=strings.Replace(t, "-", "",-1)
    return "00"+t[0:10]
}

func (c *ClusterConfig) AddWorker(ip string) *Machine {

    m :=new(Machine)
    mid :=len(c.Workers)+1
    m.Name="worker"+strconv.Itoa(mid)

    sshPortRef++
    m.SSHProxyPort=sshPortRef
    m.IP=ip
    m.MacVB=genMac()
    m.Mac=toForm(m.MacVB,":")
    c.Workers=append(c.Workers,m)
    return m

}

func (c *ClusterConfig) CreateKabnet(ip string) *Machine {
    m :=new(Machine)
    m.Name="kabnet"
    m.IP=ip
    sshPortRef++
    m.SSHProxyPort=sshPortRef
    m.MacVB=genMac()
    m.Mac=toForm(m.MacVB,":")
    c.Kabnet=m
    return m

}
func (c *ClusterConfig) CreateBootstrap(ip string) *Machine {
  m :=new(Machine)
  m.Name="bootstrap"
  sshPortRef++
  m.SSHProxyPort=sshPortRef
  m.IP=ip
  m.MacVB=genMac()
  m.Mac=toForm(m.MacVB,":")
  c.BootStrap=m
  return m
}
