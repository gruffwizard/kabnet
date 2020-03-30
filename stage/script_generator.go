package stage



import (

 "github.com/gruffwizard/kabnet/defs"
 "strings"
)

/*
Create the script element for installing packages on the Kabnet server
*/

func genPackageInstall() string {

    install := "\n#\n#Update\n#"
    install += "\nDEBIAN_FRONTEND=noninteractive apt-get update && apt-get upgrade -y\n"

    install += "\n#\n#Install\n#"

    for _, v := range defs.Kabnet.Elements() {

        desc:=v.PackageDesc()
        if desc !="" {install+="\n# "+desc}

    }

    install +="\n\napt-get install -y "+strings.Join(defs.Kabnet.PackageNames()," ")

    return install
}

/*
  Create the script element for configuring pakcages on the Kabnet server
*/
func genConfigCmds() string {

    cmds :="\n#\n#Config\n#"

    for _, v := range defs.Kabnet.Elements() {
        for  _,e := range v.Commands() {
            cmds+="\n"+e+"   #"+v.Unit()
        }
     }

     cmds+="\n"



    return cmds
}
func genStartList() string {

  start :="\n#\n#Start\n#"

  for _, v := range defs.Kabnet.Elements() {
      for  _,e := range v.StartCommands() {
          start+="\n"+e+"   #"+v.Unit()
      }
   }

    return start
}


func genSetupScript() string {

    script:="\n#\n#Copy\n#"


      for _, v := range defs.Kabnet.Elements() {
          for  _,e := range v.Files() {
            if e.ServerName!="" {
              script+="\ncp /installation/"+e.LocalName+" "+e.ServerName+" #"+v.Unit()
            }
          }
       }

    return script
}
