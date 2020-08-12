package defs



import (
  "strconv"
)


var sshPortRef int = 1200

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

    mid :=len(c.Masters)+1
    m :=newMachine("master"+strconv.Itoa(mid),ip)
    c.Masters=append(c.Masters,m)
    m.parent=c
    return m

}


func (c *ClusterConfig) AddWorker(ip string) *Machine {

    mid :=len(c.Workers)+1
    m :=newMachine("worker"+strconv.Itoa(mid),ip)
    m.parent=c
    c.Workers=append(c.Workers,m)
    return m

}

func (c *ClusterConfig) CreateKabnet(ip string) *Machine {


    m :=newMachine("kabnet",ip)
    m.parent=c
    c.Kabnet=m
    return m

}
func (c *ClusterConfig) CreateBootstrap(ip string) *Machine {

  m :=newMachine("bootstrap",ip)
  m.parent=c
  c.BootStrap=m
  return m
}
