package generate

import (
  "github.com/gruffwizard/kabnet/defs"
  "fmt"
  "strings"
  "strconv"
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

func AddBind9DNS() {

    ce:=defs.NewConfigElement("bind9")

    ce.AddPackage("bind9")
    ce.AddPackage("bind9utils")
    ce.AddPackage("resolvconf")

    ce.AddFile("bind9","/etc/default/bind9",bind9)

    ce.AddStartCmd("systemctl restart bind9")
    ce.AddStartCmd("export DRUM=\"{{.Cluster.Kabnet.Name}}.{{.Cluster.Domain}}\"")
    ce.AddStartCmd("hostnamectl set-hostname $DRUM --static")
    ce.AddStartCmd("hostname $DRUM")
    ce.AddStartCmd("ping -c 5 google.com")
    ce.AddStartCmd("ping -c 5 $DRUM")

    ce.AddCallback(callback9)
}

func callback9(ce *defs.ConfigElement,g *defs.GeneratorDefinition) {

      t:=buildBindDBTemplate(g)
      ce.AddFile("bind.db","/etc/bind/db.{{.Cluster.Domain}}",t)
      ce.AddFile("named.conf.options","/etc/bind/named.conf.options",named_conf_options)
      ce.AddFile("named.conf.local","/etc/bind/named.conf.local",named_conf_local)



      ce.AddConfigCmd("named-checkconf")
      ce.AddConfigCmd("named-checkzone {{.Cluster.Domain}} /etc/bind/db.{{.Cluster.Domain}}")

      ce.AddConfigCmd("echo 'nameserver {{.Cluster.Kabnet.IP}}' > /etc/resolvconf/resolv.conf.d/head")
      ce.AddConfigCmd("echo 'search     {{.Cluster.Domain}}' >>     /etc/resolvconf/resolv.conf.d/head")

      ce.AddStartCmd("service resolvconf restart")
}


var bind9=
`RESOLVCONF=no
OPTIONS="-u bind -4"`


var named_conf_local=

`zone "{{.Cluster.Domain}}" {
     type master;
     file "/etc/bind/db.{{.Cluster.Domain}}";
};
`

var named_conf_options=
`options {
        directory "/var/cache/bind";
        auth-nxdomain no;
        listen-on port 53 { localhost; {{.Cluster.AddressPool}}; };
        allow-query { localhost; {{.Cluster.AddressPool}}; };
        forwarders { 8.8.8.8; };
        recursion yes;
        };
`
