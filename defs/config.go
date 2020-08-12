package defs


type Config struct {
  OpenshiftVersion string `yaml:"openshift"`
  ToolsVersion string   `yaml:"tools"`
  Domain  string `yaml:"domain"`
  Cluster string `yaml:"cluster"`
  AddressPool string `yaml:"addresspool"`
  Masters int    `yaml:"masters"`
  Workers int    `yaml:"workers"`

  ImageInfo Dependencies `yaml:"deps,omitempty"`
  OpenshiftTools OpenshiftTools `yaml:"ostools,omitempty"`

}
