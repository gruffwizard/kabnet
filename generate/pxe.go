package generate

import (
  "fmt"
  "github.com/gruffwizard/kabnet/defs"

)

func init() {

    ce:=defs.NewConfigElement("pxe")

    ce.AddPackage("tftpd-hpa")
    ce.AddStartCmd("systemctl enable tftpd-hpa")

    ce.AddCallback(callbackPxe)
}


func callbackPxe(ce *defs.ConfigElement,g *defs.GeneratorConfig) {

  c:=g.Meta


  ce.AddFile("bootstrap.pxe","/var/lib/tftpboot/bootstrap.pxe",bootstrap)

  for m := 0; m < c.Masters; m++ {
      pxefile:=fmt.Sprintf("master%d.pxe",m+1)
      ce.AddFile(pxefile,"/var/lib/tftpboot/"+pxefile,master)

  }

  for w := 0; w < c.Workers; w++ {
    pxefile:=fmt.Sprintf("worker%d.pxe",w+1)
      ce.AddFile(pxefile,"/var/lib/tftpboot/"+pxefile,worker)
  }

}

var bootstrap=
`#!ipxe

set images       http://{{.Cluster.Kabnet.IP}}:8080/images
set installation http://{{.Cluster.Kabnet.IP}}:8080/installation

initrd ${images}/rhcos-4.3.0-x86_64-installer-initramfs.img
kernel  ${images}/rhcos-4.3.0-x86_64-installer-kernel -o vmware_raw ip=dhcp rd.neednet=1 initrd=${images}/rhcos-4.3.0-x86_64-installer-initramfs.img console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${images}/rhcos-4.3.0-x86_64-metal.raw.gz coreos.inst.ignition_url=${installation}/bootstrap.ign
boot
`

var worker=
`#!ipxe

set images       http://{{.Cluster.Kabnet.IP}}:8080/images
set installation http://{{.Cluster.Kabnet.IP}}:8080/installation

initrd ${images}/rhcos-4.3.0-x86_64-installer-initramfs.img
kernel  ${images}/rhcos-4.3.0-x86_64-installer-kernel -o vmware_raw ip=dhcp rd.neednet=1 initrd=${images}/rhcos-4.3.0-x86_64-installer-initramfs.img console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${images}/rhcos-4.3.0-x86_64-metal.raw.gz coreos.inst.ignition_url=${installation}/worker.ign
boot
`


var master=
`#!ipxe

set images       http://{{.Cluster.Kabnet.IP}}:8080/images
set installation http://{{.Cluster.Kabnet.IP}}:8080/installation

initrd ${images}/rhcos-4.3.0-x86_64-installer-initramfs.img
kernel  ${images}/rhcos-4.3.0-x86_64-installer-kernel -o vmware_raw ip=dhcp rd.neednet=1 initrd=${images}/rhcos-4.3.0-x86_64-installer-initramfs.img console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${images}/rhcos-4.3.0-x86_64-metal.raw.gz coreos.inst.ignition_url=${installation}/master.ign
boot
`
