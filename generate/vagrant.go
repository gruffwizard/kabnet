package generate


import (
  "github.com/gruffwizard/kabnet/defs"
)

func AddVagrantFile() {

    ce:=defs.NewConfigElement("vagrant")
    ce.AddFile("Vagrantfile","",vagrant_template)
}



var vagrant_template =
`
Vagrant.configure("2") do |config|

config.ssh.insert_key = true
config.ssh.forward_agent = true
config.ssh.private_key_path = "{{.SecretsDir}}/core_rsa"
config.ssh.username = "core"


  {{if .Options.UseVagrantDNS}}
  config.dns.tld = "{{.Cluster.Domain}}"
  {{end}}
  config.vm.define "kabnet" do |installer|

    installer.vm.provider "virtualbox" do |v|
      v.name = "kabnet"
      v.customize ["natnetwork","modify","--netname","kabnetnet","--dhcp","off"]
      v.customize ["modifyvm", :id, "--usb", "off"]
    end

    installer.vm.network "forwarded_port", guest: {{.Cluster.BootStrap.SSHProxyPort}}, host: {{.Cluster.BootStrap.SSHProxyPort}}
    {{range .Cluster.Masters}}installer.vm.network "forwarded_port", guest: {{.SSHProxyPort}}, host: {{.SSHProxyPort}}
    {{end}}
    {{range .Cluster.Workers}}installer.vm.network "forwarded_port", guest: {{.SSHProxyPort}}, host: {{.SSHProxyPort}}
    {{end}}

    installer.vm.box = "ubuntu/xenial64"
    installer.vm.network :private_network, ip: '{{.Gateway}}'
    installer.vm.network :private_network, ip: "{{.Cluster.Kabnet.IP}}" , :name => 'vboxnet0', virtualbox__intnet: "kabnetnet"
    installer.vm.synced_folder "{{.Installation}}", "/installation"
    installer.vm.synced_folder "{{.Installer}}/{{.Meta.ToolsVersion}}", "/installer"
    installer.vm.synced_folder "{{.ImageDir}}", "/images"
    installer.vm.provision "shell", path: "kabnet.sh"
    installer.vm.provision "shell", path: "kabnet-check.sh"

    installer.vm.hostname = "{{.Cluster.Kabnet.Name}}"

    installer.ssh.insert_key = false
    installer.ssh.forward_agent = true
    installer.ssh.private_key_path = nil
    installer.ssh.username = "vagrant"
  end



  config.vm.define "bootstrap" do |client|
      client.vm.box = "osuosl/pxe"
      client.vm.network :private_network, ip: "{{.Cluster.BootStrap.IP}}" , :mac => "{{.Cluster.BootStrap.MacVB}}", virtualbox__intnet: "kabnetnet",:name => 'vboxnet0', :adapter => 1


      client.vm.provider "virtualbox" do |v|
        v.name = "bootstrap"
          v.customize ["modifyvm", :id, "--nattftpserver1", "{{.Cluster.Kabnet.IP}}"]
          v.customize ["modifyvm", :id, "--nattftpfile1"  , "{{.Cluster.BootStrap.Name}}.pxe"]
          v.customize ["storageattach",:id,"--storagectl", "IDE",
                       "--port","1","--device","0","--type","dvddrive","--medium","{{.ImageDir}}/ipxe.iso"]
          v.customize ["modifyvm",:id,"--boot1","disk","--boot2","dvd","--boot3","floppy","--boot4","net"]
          v.customize ["modifyvm", :id, "--usb", "off"]

          v.memory = 8192


      end

      client.ssh.username = "core"
      client.ssh.insert_key = false
      client.ssh.host = "{{.Gateway}}"
      client.ssh.port = {{.Cluster.BootStrap.SSHProxyPort}}
      client.vm.hostname = "{{.Cluster.BootStrap.Name}}"
    end


    {{range .Cluster.Masters}}
    config.vm.define "{{.Name}}" do |m|
      m.vm.box = "osuosl/pxe"


      m.vm.network :private_network, ip: "{{.IP}}" ,auto_config: false, :mac => "{{.MacVB}}", :name => 'vboxnet0', :adapter => 1, virtualbox__intnet: "kabnetnet"



      m.vm.provider "virtualbox" do |v|
        v.name = "{{.Name}}"
        v.customize ["modifyvm", :id, "--nattftpserver1", "{{$.Cluster.Kabnet.IP}}"]
        v.customize ["modifyvm", :id, "--nattftpfile1"  , "{{.Name}}.pxe"]
        v.customize ["storageattach",:id,"--storagectl", "IDE",
        "--port","1","--device","0","--type","dvddrive","--medium","{{$.ImageDir}}/ipxe.iso"]
        v.customize ["modifyvm",:id,"--boot1","disk","--boot2","dvd","--boot3","floppy","--boot4","net"]
        v.customize ["modifyvm", :id, "--usb", "off"]

        v.memory = 8192

      end

      m.ssh.username = "core"
      m.ssh.insert_key = false
      m.ssh.host = "{{$.Gateway}}"
      m.ssh.port = {{.SSHProxyPort}}

      m.vm.hostname = "{{.Name}}.{{$.Cluster.Cluster}}"
    end
   {{end}}

   {{range .Cluster.Workers}}
   config.vm.define "{{.Name}}" do |m|
     m.vm.box = "osuosl/pxe"
     m.vm.network :private_network, ip: "{{.IP}}" ,auto_config: false, :mac => "{{.MacVB}}",  :adapter => 1, virtualbox__intnet: "kabnetnet"
     m.vm.hostname = "{{.Name}}.{{$.Cluster.Cluster}}"


     m.vm.synced_folder ".", "/vagrant", disabled: true
     m.vm.provider "virtualbox" do |v|
       v.name = "{{.Name}}"
       v.customize ["modifyvm", :id, "--nattftpserver1", "{{$.Cluster.Kabnet.IP}}"]
       v.customize ["modifyvm", :id, "--nattftpfile1"  , "{{.Name}}.pxe"]
       v.customize ["storageattach",:id,"--storagectl", "IDE",
       "--port","1","--device","0","--type","dvddrive","--medium","{{$.ImageDir}}/ipxe.iso"]
       v.customize ["modifyvm",:id,"--boot1","disk","--boot2","dvd","--boot3","floppy","--boot4","net"]
       v.customize ["modifyvm", :id, "--usb", "off"]

       v.memory = 4096

     end

     m.ssh.username = "core"
     m.ssh.insert_key = false
     m.ssh.host = "{{$.Gateway}}"
     m.ssh.port = {{.SSHProxyPort}}


   end
  {{end}}


end
`
