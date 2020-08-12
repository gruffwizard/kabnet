package stage




import (
  "os"
  "github.com/gruffwizard/kabnet/defs"
  "github.com/gruffwizard/kabnet/util"
  "github.com/apparentlymart/go-cidr/cidr"
  "log"
  "net"
  "strings"
)

/*
  Creates install-config.yaml and then drives openshift installer
  to generate details configs.

*/
func Generate(g *defs.GeneratorDefinition) {

  checkOptions(g)

  // delete installation directory ...
  

  util.SpeculativeExecute("VBoxManage","natnetwork","add","--netname","kabnetnet","--network","vboxnet1")


  ip,netaddr,err := net.ParseCIDR(g.Cluster.AddressPool)
  if err != nil { log.Fatal(err) }

  g.Cluster.SubnetIP=ip.String()

  _,to := cidr.AddressRange(netaddr)

  g.Cluster.AddressRange=[]string{g.Cluster.BootStrap.IP,to.String()}
  g.Cluster.NetMask=net.IP(netaddr.Mask).String()

  log.Print(g.Cluster.AddressRange)
  log.Print(g.Cluster.NetMask)

  g.Secret=util.LoadFile(g.SecretsDir+"/pull-secret.txt")
  g.Key=util.LoadFile(g.SecretsDir+"/core_rsa.pub")



  kabnetSetup(g)

  //out, err := exec.Command("vagrant","-v").Output()
  //if err != nil { log.Fatal(err)}
  //util.Info("vagrant version: %s", out)

util.Section("Starting install server using vagrant")
os.Chdir(g.Installer)
//util.Execute("vagrant", "up")
//util.Info("Removing installer machine")
//util.Execute("vagrant", "destroy","-f")


}


func checkOptions(g *defs.GeneratorDefinition) {

  info,_:=util.SpeculativeExecute("vagrant","plugin","list")

  g.Options.UseVagrantDNS=strings.Contains(info,"vagrant-dns")

}
