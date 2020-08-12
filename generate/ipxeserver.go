package generate

import (
  "fmt"
  "github.com/gruffwizard/kabnet/defs"

)

func AddIPXEServer() {

    ce:=defs.NewConfigElement("ipxe")
    ce.AddPackage("ipxe")
    ce.AddConfigCmd("mkdir -p /var/lib/tftpboot")
    ce.AddConfigCmd("cp /usr/lib/ipxe/{undionly.kpxe,ipxe.efi} /var/lib/tftpboot")
    ce.AddCallback(callbackPxe)
}


func callbackPxe(ce *defs.ConfigElement,g *defs.GeneratorDefinition) {

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

set installation http://{{.Cluster.Kabnet.IP}}:8080/installation/bootstrap.ign
set ramfs        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Initramfs}}
set kernel       http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Kernel}}
set metal        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Metal}}

initrd  ${ramfs}
kernel  ${kernel} -o vmware_raw ip=dhcp rd.neednet=1 initrd=${ramfs} console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${metal} coreos.inst.ignition_url=${installation}
boot
`

var worker=
`#!ipxe

set installation http://{{.Cluster.Kabnet.IP}}:8080/installation/worker.ign
set ramfs        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Initramfs}}
set kernel       http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Kernel}}
set metal        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Metal}}

initrd  ${ramfs}
kernel  ${kernel} -o vmware_raw ip=dhcp rd.neednet=1 initrd=${ramfs} console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${metal} coreos.inst.ignition_url=${installation}
boot
`


var master=
`#!ipxe

set installation http://{{.Cluster.Kabnet.IP}}:8080/installation/master.ign
set ramfs        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Initramfs}}
set kernel       http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Kernel}}
set metal        http://{{.Cluster.Kabnet.IP}}:8080/images/{{.Meta.OpenshiftVersion}}/{{.Meta.ImageInfo.Metal}}

initrd  ${ramfs}
kernel  ${kernel} -o vmware_raw ip=dhcp rd.neednet=1 initrd=${ramfs} console=tty0 console=ttyS0 coreos.inst=yes coreos.inst.install_dev=sda coreos.inst.image_url=${metal} coreos.inst.ignition_url=${installation}
boot
`
