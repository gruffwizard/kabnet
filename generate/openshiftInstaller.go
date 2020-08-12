package generate




import (

  "github.com/gruffwizard/kabnet/defs"

)


func AddOpenShiftInstaller() {

    ce:=defs.NewConfigElement("openshift-installer")
    ce.AddFile("install-config.yaml","",template1)
    ce.AddConfigCmd("chmod +x /installer/openshift-install")
    ce.AddConfigCmd("/installer/openshift-install create manifests --dir=/installation")
    ce.AddConfigCmd("sed -i 's/mastersSchedulable: true/mastersSchedulable: false/g' /installation/manifests/cluster-scheduler-02-config.yml")
    ce.AddConfigCmd("/installer/openshift-install create ignition-configs --dir=/installation")
    ce.AddConfigCmd("chmod -R a+rw /installation")
}

//var check_script=`openshift-install gather bootstrap --help `

var template1=
`apiVersion: v1
baseDomain: {{.Cluster.Domain}}
compute:
- hyperthreading: Enabled
  name: worker
  replicas: 0
controlPlane:
  hyperthreading: Enabled
  name: master
  replicas: {{.Meta.Masters}}
metadata:
  name: {{.Cluster.Cluster}}
networking:
  clusterNetwork:
  - cidr: 10.128.0.0/14
    hostPrefix: 23
  networkType: OpenShiftSDN
  serviceNetwork:
  - 172.30.0.0/16
platform:
  none: {}
fips: false
pullSecret: '{{.Secret}}'
sshKey: '{{.Key}}'`
