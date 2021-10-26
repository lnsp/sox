package cloudconfig

type User struct {
	Name           string `yaml:"name"`
	Sudo           string `yaml:"sudo"`
	Groups         string `yaml:"groups"`
	Shell          string `yaml:"shell"`
	Home           string `yaml:"home"`
	LockPasswd     bool   `yaml:"lock_passwd"`
	AuthorizedKeys string `yaml:"ssh_authorized_keys"`
}

type Chpasswd struct {
	Expire bool     `yaml:"expire"`
	List   []string `yaml:"list"`
}

type SetHostnameModule struct {
	Hostname string `yaml:"hostname"`
	FQDN     string `yaml:"fqdn"`
}

type ManageEtcHostsModule struct {
	ManageEtcHosts bool `yaml:"manage_etc_hosts"`
}

type UsersModule struct {
	Users []User `yaml:"users"`
}

type SetPasswordsModule struct {
	SSHPwAuth bool     `yaml:"ssh_pwauth"`
	Chpasswd  Chpasswd `yaml:"chpasswd"`
}

type SSHModule struct {
	DisableRoot bool `yaml:"disable_root"`
}

type PackageModule struct {
	Packages       []string `yaml:"packages"`
	PackageUpdate  bool     `yaml:"package_update"`
	PackageUpgrade bool     `yaml:"package_upgrade"`
}

type CloudConfig struct {
	SetHostnameModule
	ManageEtcHostsModule
	UsersModule
	SetPasswordsModule
	PackageModule
}

type NetworkNameservers struct {
	Addresses     []string `yaml:"addresses"`
	SearchDomains []string `yaml:"search"`
}

type NetworkEthernet struct {
	DHCPv4      bool               `yaml:"dhcp4"`
	Addresses   []string           `yaml:"addresses"`
	GatewayIPv4 string             `yaml:"gateway4"`
	Nameservers NetworkNameservers `yaml:"nameservers"`
}

type NetworkConfig struct {
	Version   int                        `yaml:"version"`
	Ethernets map[string]NetworkEthernet `yaml:"ethernets"`
}
