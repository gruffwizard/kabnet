package stage


import (
  "github.com/gruffwizard/kabnet/defs"
  "fmt"
)

// Iptables needed to allow the nodes to communicate with the outside world

func init() {

    ce:=defs.NewConfigElement("iptables")
    ce.AddCallback(callbackIPTables)

}

func callbackIPTables(ce *defs.ConfigElement,g *defs.GeneratorDefinition) {

    ce.AddConfigCmd("echo 1 > /proc/sys/net/ipv4/ip_forward")

    ce.AddConfigCmd("export LAN=`ip route | grep {{.Cluster.Kabnet.IP}} | cut -d' ' -f3`")
    ce.AddConfigCmd("export WAN=`ip route | grep default | cut -d' ' -f5`")

    ce.AddConfigCmd("iptables --table nat --append POSTROUTING --out-interface $WAN -j MASQUERADE")
    ce.AddConfigCmd("iptables --append FORWARD --in-interface $LAN  -j ACCEPT")

    ce.AddConfigCmd(genNat(g.Cluster.BootStrap))

    for _, v := range g.Cluster.Masters {
        ce.AddConfigCmd(genNat(v))
    }

    for _, v := range g.Cluster.Workers {
        ce.AddConfigCmd(genNat(v))
    }





}

func genNat(m *defs.Machine) string {

  return fmt.Sprintf("iptables -t nat -A PREROUTING -p tcp -d {{.Gateway}} --dport %d -j DNAT --to-destination %s:22",
    m.SSHProxyPort,m.IP)

}
