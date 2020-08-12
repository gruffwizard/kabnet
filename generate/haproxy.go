package generate  

import (
  "github.com/gruffwizard/kabnet/defs"
)

func AddHAProxy() {

    ce:=defs.NewConfigElement("haproxy")
    ce.AddFile("haproxy.conf","/etc/haproxy/haproxy.cfg",ha_config)
    ce.AddPackage("haproxy")
    ce.AddStartCmd("service haproxy restart")

}


var ha_config string=
`
global
  daemon
  log 127.0.0.1 local0
  log 127.0.0.1 local1 notice
  maxconn 4096
  tune.ssl.default-dh-param 2048

defaults
  log               global
  retries           3
  maxconn           2000
  timeout connect   5s
  timeout client    50s
  timeout server    50s
  mode              tcp
  option            tcplog

listen stats # Define a listen section called "stats"
    bind :9090 # Listen on localhost:9090
    mode http
    stats enable  # Enable stats page
    stats hide-version  # Hide HAProxy version
    stats realm Haproxy\ Statistics  # Title text for popup window
    stats uri /haproxy_stats  # Stats URI
    stats auth admin:admin


frontend openshift-api-server
    bind *:6443
    default_backend openshift-api-server
    mode              tcp


backend openshift-api-server
    balance source
    mode              tcp
    option ssl-hello-chk

    server {{.Cluster.BootStrap.Name}} {{.Cluster.BootStrap.IP}}:6443 check
{{range .Cluster.Masters}}
    server {{.Name}} {{.IP}}:6443 check
{{end}}

frontend machine-config-server
    bind *:22623
    default_backend machine-config-server
    mode              tcp

backend machine-config-server
    balance source
    mode              tcp
    option ssl-hello-chk

    server {{.Cluster.BootStrap.Name}} {{.Cluster.BootStrap.IP}}:22623 check
{{range .Cluster.Masters}}
    server {{.Name}} {{.IP}}:22623 check
{{end}}


frontend ingress-http
    bind *:80
    default_backend ingress-http
    mode              tcp


backend ingress-http
    balance source
    mode              tcp
    option ssl-hello-chk

  {{range .Cluster.Workers}}
    server {{.Name}} {{.IP}}:80 check
  {{end}}

frontend ingress-https
    bind *:443
    default_backend ingress-https
    mode              tcp


backend ingress-https
    balance source
    mode              tcp
    option ssl-hello-chk

{{range .Cluster.Workers}}
    server {{.Name}} {{.IP}}:443 check
{{end}}


        `
