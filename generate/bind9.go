package generate

import (
  "github.com/gruffwizard/kabnet/defs"
)


func init() {

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

func callback9(ce *defs.ConfigElement,g *defs.GeneratorConfig) {

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
