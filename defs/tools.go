package defs


type Dependencies struct {
    BaseURL string
    Initramfs string
    Kernel string
    Metal  string

}


func (d *Dependencies) FileNames() []string {

  return []string{ d.Initramfs,d.Kernel,d.Metal}
}

type OpenshiftTools struct {
    BaseURL string
    Client string
    Installer string

}

func (d *OpenshiftTools) FileNames() []string {

  return []string{ d.Client,d.Installer}
}
