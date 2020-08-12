package generate

import (
    "github.com/gruffwizard/kabnet/defs"
)

func AddDnsmasqServer() {

  ce:=defs.NewConfigElement("dnsmasq")

  ce.AddPackage("dnsmasq")

  ce.AddConfigCmd("mkdir -p /var/lib/tftpboot")
  ce.AddConfigCmd("chown dnsmasq:nogroup /var/lib/tftpboot/*")

  ce.AddConfigCmd("mv /etc/resolv.conf  /etc/resolv.conf.orig")


  ce.AddStartCmd("systemctl enable dnsmasq")
  ce.AddStartCmd("systemctl start  dnsmasq")

  ce.AddStartCmd("systemctl restart systemd-resolved")
  ce.AddStartCmd("systemctl restart dnsmasq")


  ce.AddFile("resolved.conf","/etc/systemd/resolved.conf",resolved_conf)
  ce.AddFile("common.conf","/etc/dnsmasq.d/common.conf",common_conf)

  ce.AddFile("dev.conf","/etc/dnsmasq.d/dev.conf",dev_conf)

}

var dev_conf string=
`
{{range .Cluster.Masters}}
dhcp-host={{.MacAddrFmt ":"}},{{.IP}}
address=/{{.Fqdn}}/{{.IP}}
ptr-record={{.ReverseIP}}.in-addr.arpa,{{.Fqdn}}
dhcp-boot={{.Name}}.pxe

{{end}}

{{range .Cluster.Workers}}
dhcp-host={{.MacAddrFmt ":"}},{{.IP}}
address=/{{.Fqdn}}/{{.IP}}
ptr-record={{.ReverseIP}}.in-addr.arpa,{{.Fqdn}}
dhcp-boot={{.Name}}.pxe
{{end}}


dhcp-host={{.Cluster.BootStrap.MacAddrFmt ":"}},{{.Cluster.BootStrap.IP}}
address=/{{.Cluster.BootStrap.Fqdn}}/{{.Cluster.BootStrap.IP}}
ptr-record={{.Cluster.BootStrap.ReverseIP}}.in-addr.arpa,{{.Cluster.BootStrap.Fqdn}}
dhcp-boot=bootstrap.pxe

`


var resolved_conf string=
`[Resolve]
DNS=127.0.0.1`

var common_conf string =
`# DHCP
dhcp-option=option:router,192.168.50.1
dhcp-option=option:dns-server,192.168.50.1
dhcp-range=192.168.50.100,192.168.50.254,12h
# forward, use original DNS server
server=8.8.8.8
server=8.8.4.4
enable-tftp
tftp-root=/var/lib/tftpboot
tftp-secure
# Legacy PXE
dhcp-match=set:bios,option:client-arch,0
dhcp-boot=tag:bios,undionly.kpxe
# UEFI
dhcp-match=set:efi32,option:client-arch,6
dhcp-boot=tag:efi32,ipxe.efi
dhcp-match=set:efibc,option:client-arch,7
dhcp-boot=tag:efibc,ipxe.efi
dhcp-match=set:efi64,option:client-arch,9
dhcp-boot=tag:efi64,ipxe.efi
dhcp-userclass=set:ipxe,iPXE
`
