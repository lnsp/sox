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

type CloudConfig struct {
	// SetHostnameModule
	Hostname string `yaml:"hostname"`
	FQDN     string `yaml:"fqdn"`

	// ManageEtcHostsModule
	ManageEtcHosts bool `yaml:"manage_etc_hosts"`

	// UsersModule
	Users []User `yaml:"users"`

	// SetPasswordsModule
	SSHPwAuth bool     `yaml:"ssh_pwauth"`
	Chpasswd  Chpasswd `yaml:"chpasswd"`

	// PackageModule
	Packages       []string `yaml:"packages"`
	PackageUpdate  bool     `yaml:"package_update"`
	PackageUpgrade bool     `yaml:"package_upgrade"`

	// SSHModule
	DisableRoot bool `yaml:"disable_root"`

	// ResolvConfModule
	ManageResolvConf bool       `yaml:"manage_resolv_conf"`
	ResolvConf       ResolvConf `yaml:"resolv_conf"`

	// WriteFilesModule
	WriteFiles []WriteFile `yaml:"write_files"`
}

type WriteFile struct {
	Content string `yaml:"content"`
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

type ResolvConfOptions struct {
	Rotate  bool `yaml:"rotate"`
	Timeout int  `yaml:"timeout"`
}

type ResolvConf struct {
	Nameservers   []string          `yaml:"nameservers"`
	SearchDomains []string          `yaml:"searchdomains"`
	Domain        string            `yaml:"domain"`
	Options       ResolvConfOptions `yaml:"options"`
}
