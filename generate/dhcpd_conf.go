package generate

import (
  "github.com/gruffwizard/kabnet/defs"
)

func init() {

    ce:=defs.NewConfigElement("dhcpd")
    ce.AddFile("dhcp.conf","/etc/dhcp/dhcpd.conf",dhcpd_conf_template)

    ce.AddPackage("isc-dhcp-server")

    ce.AddConfigCmd("export LAN=`ip route | grep {{.Cluster.Kabnet.IP}} | cut -d' ' -f3`")
    ce.AddConfigCmd("export WAN=`ip route | grep default | cut -d' ' -f5`")
    ce.AddConfigCmd("echo \"INTERFACES=$LAN\" > /etc/default/isc-dhcp-server")
    ce.AddConfigCmd("touch /var/lib/dhcp/dhcpd.leases")
    ce.AddStartCmd("dhcpd -t")
    ce.AddStartCmd("systemctl restart isc-dhcp-server.service")


}

var dhcpd_conf_template =
`
#configuration file for ISC dhcpd
ddns-update-style none;
option domain-name "{{.Cluster.Domain}}";
#option bootfile-name "pxelinux1.0";
option domain-name-servers {{.Cluster.Kabnet.IP}};
default-lease-time 3600;
max-lease-time 7200;
authoritative;

subnet {{.Cluster.SubnetIP}} netmask 255.255.255.0 {
    option routers {{.Cluster.Kabnet.IP}};
    option domain-name-servers {{.Cluster.Kabnet.IP}};
    option subnet-mask {{.Cluster.NetMask}};
    option domain-name "{{.Cluster.Domain}}";
    range {{index .Cluster.AddressRange 0}} {{index .Cluster.AddressRange 1}};
}

host {{.Cluster.Kabnet.Name}} {
  hardware ethernet {{.Cluster.Kabnet.Mac}};
  fixed-address {{.Cluster.Kabnet.IP}};
}

host {{.Cluster.BootStrap.Name}} {
  hardware ethernet    {{.Cluster.BootStrap.Mac}};
  option bootfile-name "{{.Cluster.BootStrap.Name}}.pxe";
  fixed-address        {{.Cluster.BootStrap.IP}};
}

{{range .Cluster.Masters}}
host {{.Name}} {
  hardware ethernet {{.Mac}};
  fixed-address {{.IP}};
  option bootfile-name "{{.Name}}.pxe";
}{{end}}
{{range .Cluster.Workers}}
host {{.Name}} {
  hardware ethernet {{.Mac}};
  fixed-address {{.IP}};
  option bootfile-name "{{.Name}}.pxe";
}{{end}}

`
