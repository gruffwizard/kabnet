package generate

import (
    "github.com/gruffwizard/kabnet/defs"
    "fmt"
    "github.com/gruffwizard/kabnet/util"
    "github.com/apparentlymart/go-cidr/cidr"
    "log"
    "net"
    "os"

)

func Generate(g *defs.GeneratorDefinition) {

      fmt.Println("running generate")

      setup(g)

      AddApacheServer()
      AddHAProxy()
      AddIPTables()

      AddVagrantFile()
      AddOpenShiftInstaller()

      //AddISCDhcpServer()
      //AddBind9DNS()
      //AddTFTPServerHPA()

      AddDnsmasqServer()

      AddIPXEServer()


      kabnetSetup(g)

    util.Section("Starting install server using vagrant")
    os.Chdir(g.Installer)

}


func setup(g *defs.GeneratorDefinition) {


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

}
