package stage

import (
  "fmt"
  "strings"
  "strconv"
    "github.com/gruffwizard/kabnet/defs"
)

func type1Fwd(name string,ipaddr string) string {
  return fmt.Sprintf("%-25s     IN  A  %s\n",name,ipaddr)
}
func type2Rev(name string,ipaddr string) string {

    lastOctetA := strings.Split(ipaddr,".")
    lastOctet  :=  lastOctetA[len(lastOctetA)-1]
    return fmt.Sprintf("                          %3s IN PTR  %s\n",lastOctet,name)
}

func genFwdMachines(machines []*defs.Machine ,cluster string) string {
  results:=""
  for w := 0; w < len(machines); w++ {
      m:=machines[w]
      results+=type1Fwd(m.Name+"."+cluster,m.IP)
  }
  return results
}

func genRevMachines(machines []*defs.Machine ,cluster string) string {
  results:=""
  for w := 0; w < len(machines); w++ {
      m:=machines[w]
      results+=type2Rev(m.Name+"."+cluster,m.IP)
  }
  return results
}

func genForwardRecords(g *defs.GeneratorDefinition) string {

  // add real machines
  r:="; Forward references\n"
  r+= genFwdMachines(g.Cluster.Masters,g.Cluster.Cluster)
  r+= genFwdMachines(g.Cluster.Workers,g.Cluster.Cluster)
  // add aliases
  r+= type1Fwd(g.Cluster.Kabnet.Name,g.Cluster.Kabnet.IP)
  r+= type1Fwd("ns1",g.Cluster.Kabnet.IP)
  r+= type1Fwd("smtp",g.Cluster.Kabnet.IP)
  r+= type1Fwd("api."+g.Cluster.Cluster,g.Cluster.Kabnet.IP)
  r+= type1Fwd("api-int."+g.Cluster.Cluster,g.Cluster.Kabnet.IP)

  // add etc servers (on masters )
  masters:=g.Cluster.Masters
  for w := 0; w < len(masters); w++ {
    name:="etcd-"+strconv.Itoa(w)+"."+g.Cluster.Cluster
    r+= type1Fwd(name,masters[w].IP)
  }

  return r
}

func genReverseRecords(g *defs.GeneratorDefinition) string {

  // add real machines
  r:="; Reverse references\n"
  r+= genRevMachines(g.Cluster.Masters,g.Cluster.Cluster)
  r+= genRevMachines(g.Cluster.Workers,g.Cluster.Cluster)
  // add aliases
  r+= type2Rev(g.Cluster.Kabnet.Name,g.Cluster.Kabnet.IP)
  r+= type2Rev("ns1",g.Cluster.Kabnet.IP)
  r+= type2Rev("smtp",g.Cluster.Kabnet.IP)
  r+= type2Rev("api."+g.Cluster.Cluster,g.Cluster.Kabnet.IP)
  r+= type2Rev("api-int."+g.Cluster.Cluster,g.Cluster.Kabnet.IP)

  // add etc servers (on masters )
  masters:=g.Cluster.Masters
  for w := 0; w < len(masters); w++ {
    name:="etcd-"+strconv.Itoa(w)+"."+g.Cluster.Cluster
    r+= type2Rev(name,masters[w].IP)
  }

  return r
}

func genEtcdRecords(g *defs.GeneratorDefinition) string {
  r:="; ETCD entries\n"
  masters:=g.Cluster.Masters
  for w := 0; w < len(masters); w++ {

r+=fmt.Sprintf("_etcd-server-ssl._tcp.%s 86400 IN SRV 0 10 2380 etcd-%d.%s.%s.\n",
  g.Cluster.Cluster,
  w,g.Cluster.Cluster,
  g.Cluster.Domain)
r+=fmt.Sprintf("_etcd-client-ssl._tcp.%s 86400 IN SRV 0 10 2379 etcd-%d.%s.%s.\n",
  g.Cluster.Cluster,
  w,g.Cluster.Cluster,
  g.Cluster.Domain)
}

return r
}

func buildBindDBTemplate(g *defs.GeneratorDefinition) string {

  text:=named_zone_header

  text+=genForwardRecords(g)

  text+=genReverseRecords(g)

  text+=genEtcdRecords(g)


  return text
}




var named_zone_header=
`$TTL 1W
@   IN  SOA     {{.Cluster.Domain}} root.{{.Cluster.Domain}} (
2019070700  ; serial
3H          ; refresh (3 hours)
30M         ; retry (30 minutes)
2W          ; expiry (2 weeks)
1W )        ; minimum (1 week)
; necessary dns references
                              IN     NS      ns1.{{.Cluster.Domain}}.
                              IN     MX 10   smtp.{{.Cluster.Domain}}.
*.apps.kab                    IN     A       {{.Cluster.Kabnet.IP}}
`
