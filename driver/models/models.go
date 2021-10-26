package models

type Machine struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	IPAddress string

	ImageID  string
	Image    Image
	SSHKeyID string
	SSHKey   SSHKey

	Specs Specs `gorm:"embedded"`
}

type Specs struct {
	// vCPU count.
	CPUs int64
	// Memory size in MB.
	Memory int64
	// Disk size in MB.
	Disk int64
}

type Image struct {
	ID   string `gorm:"primaryKey"`
	Name string
	OS   string
	Path string
}

type SSHKey struct {
	ID     string `gorm:"primaryKey"`
	Name   string
	Pubkey string
}
