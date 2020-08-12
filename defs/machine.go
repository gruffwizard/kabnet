package defs


import (
  "strings"
)
type Machine struct {

    Name string `yaml:"name"`
    IP  string   `yaml:"ip"`
    Mac string `yaml:"mac"`
    MacVB string `yaml:"macvb,omitempty"`
    SSHProxyPort int `yaml:"sshproxyport"`
    parent *ClusterConfig


}


func newMachine(name string,ip string) (*Machine) {

  m:=new (Machine)

  m.Name=name
  m.IP=ip
  sshPortRef++
  m.SSHProxyPort=sshPortRef
  m.MacVB=genMac()
  m.Mac=toForm(m.MacVB,":")

  return m
}

func (m *Machine) ReverseIP() string {

    ip:=strings.Split(m.IP,".")
    var rip []string

     for i := len(ip)-1; i >= 0; i-- {
      rip=append(rip,ip[i])
    }

    return strings.Join(rip,".")
}

func (m *Machine) Fqdn() string {

    return m.Name+"."+m.parent.Cluster+"."+m.parent.Domain
}

func (m *Machine) MacAddrFmt(fmt string) string {

    return toForm(m.Mac,fmt)
}

func toForm(mac string,s string) string {

    r := mac[0:2] + s + mac[2:4] + s + mac[4:6] + s + mac[6:8] + s + mac[8:10] + s + mac[10:12]
    return r

}
