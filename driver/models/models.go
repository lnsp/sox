package models

import (
	"fmt"

	"gorm.io/gorm"
)

type MachineState int64

const (
	StateUnknown MachineState = iota
	StateCreated
	StateStopped
	StateRunning
	StateCrashed
)

type Machine struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`

	ImageID string
	Image   Image

	SSHKeys []SSHKey `gorm:"many2many:machine_ssh_keys"`

	Specs             Specs              `gorm:"embedded"`
	NetworkInterfaces []NetworkInterface `gorm:"foreignkey:machine_id"`
}

type Specs struct {
	// vCPU count.
	CPUs int64
	// Memory size in MB.
	Memory int64
	// Disk size in GB.
	Disk int64
}

func (s Specs) String() string {
	return fmt.Sprintf("%d vCPUs, %dMiB memory, %dGiB Disk Space", s.CPUs, s.Memory, s.Disk)
}

type Image struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`
	OS   string
	Path string
}

type SSHKey struct {
	ID     string `gorm:"primaryKey"`
	Name   string `gorm:"uniqueIndex"`
	Pubkey string
}

type Activity struct {
	gorm.Model

	Type    string
	Subject string
}

type NetworkInterface struct {
	ID int64 `gorm:"primaryKey"`

	MachineID string
	NetworkID string `gorm:"index:idx_ip4,unique"`
	Network   Network

	IPv4 string `gorm:"index:idx_ip4,unique"`
	IPv6 string

	HwAddr string
}

type Network struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex"`

	IPv4 NetworkSpec `gorm:"embedded;embeddedPrefix:ipv4_"`
	IPv6 NetworkSpec `gorm:"embedded;embeddedPrefix:ipv6_"`

	Nameservers   string
	SearchDomains string
}

type NetworkSpec struct {
	Subnet  string
	Gateway string
}
