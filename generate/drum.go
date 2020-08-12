package generate


import (

  "github.com/gruffwizard/kabnet/util"
    "github.com/gruffwizard/kabnet/defs"

)

var kabnet_check string=
`
# Kabnet
dig {{.Cluster.Kabnet.Name}}.{{.Cluster.Domain}} +short
dig -x {{.Cluster.Kabnet.IP}} +short

# Bootstrap
dig {{.Cluster.BootStrap.Name}}.{{.Cluster.Domain}} +short
dig -x {{.Cluster.BootStrap.IP}} +short

# Masters
{{range .Cluster.Masters}}
dig {{.Name}}.{{$.Cluster.Cluster}}.{{$.Cluster.Domain}} +short
dig -x {{.IP}} +short
{{end}}

# Workers
{{range .Cluster.Workers}}
dig {{.Name}}.{{$.Cluster.Cluster}}.{{$.Cluster.Domain}} +short
dig -x {{.IP}} +short
{{end}}

# API
dig api.{{.Cluster.Cluster}}.{{.Cluster.Domain}} +short
dig api-int.{{.Cluster.Cluster}}.{{.Cluster.Domain}} +short

# Wildcard
dig *.apps.{{.Cluster.Cluster}}.{{.Cluster.Domain}} +short

# ETCD
dig _etcd-server-ssl._tcp.{{.Cluster.Cluster}}.{{.Cluster.Domain}} SRV +short
dig _etcd-client-ssl._tcp.{{.Cluster.Cluster}}.{{.Cluster.Domain}} SRV +short

`

func kabnetSetup(g *defs.GeneratorDefinition) {


  g.Installation=util.RecreateDir(g.RootDir,"installation")
  // write all the data files..

  util.Info("Creating Vagrant file for installer")

  // run callbacks for installers to do configs..
  defs.Kabnet.Callback(g)

  // firstly write all the script files to the installation directory
  install:=""

  for _, v := range defs.Kabnet.Files() {

      util.WriteTemplate(g,g.Installation+"/"+v.LocalName,v.Content)
  }


  install+=genPackageInstall()
  install+=genSetupScript()
  install+=genConfigCmds()
  install+=genStartList()

  util.WriteTemplate(g,g.Installation+"/kabnet.sh",install)
  util.WriteTemplate(g,g.Installation+"/kabnet-check.sh",kabnet_check)



}
