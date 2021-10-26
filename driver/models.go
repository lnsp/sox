package driver

type Machine struct {
	ID        string `gorm:"uniqueIndex"`
	Name      string
	IPAddress string
}

type Image struct {
	ID   string `gorm:"uniqueIndex"`
	Name string
	Path string
}

type SSHKey struct {
	ID     string `gorm:"uniqueIndex"`
	Pubkey string
}
