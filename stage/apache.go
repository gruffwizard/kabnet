package stage


import (
  "github.com/gruffwizard/kabnet/defs"
)

func init() {

    ce:=defs.NewConfigElement("apache")

    ce.AddPackage("apache2")

    ce.AddFile("ports.conf","/etc/apache2/ports.conf",ports)
    ce.AddConfigCmd("sed -i 's/VirtualHost *:80/VirtualHost *:8080/g' /etc/apache2/sites-enabled/000-default.conf")
    ce.AddConfigCmd("ln -s /installation /var/www/html/installation")
    ce.AddConfigCmd("chmod a+rw          /var/www/html/installation")
    ce.AddConfigCmd("ln -s /images       /var/www/html/images")
    ce.AddConfigCmd("chmod a+rw          /var/www/html/images")

    ce.AddStartCmd("service apache2 restart")


}


var ports string =`Listen 8080
      <IfModule ssl_module>
        Listen 9443
      </IfModule>
      <IfModule mod_gnutls.c>
        Listen 9443
      </IfModule>
`
