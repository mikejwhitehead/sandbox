package plugin

func init() {
	
}

type VCenterPluginConfig
// VCenterPlugin struct
type VCenterPlugin struct {
	VCenters map[string]VCenter

}

// VCenter struct
type VCenter struct {

	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Insecure bool   `yaml:"insecure"`

	Client *vim25.Client
	REST   *rest.Client

	logger *logrus.Logger
}