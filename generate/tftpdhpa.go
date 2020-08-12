package generate

import (
  "github.com/gruffwizard/kabnet/defs"

)

func AddTFTPServerHPA() {

    ce:=defs.NewConfigElement("tftpdhpa")
    ce.AddPackage("tftpd-hpa")
    ce.AddStartCmd("systemctl enable tftpd-hpa")

  }
