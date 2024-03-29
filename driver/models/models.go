package models

import (
	"fmt"
	"path/filepath"

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

	User    string
	SSHKeys []SSHKey `gorm:"many2many:machine_ssh_keys"`

	Specs             Specs              `gorm:"embedded"`
	NetworkInterfaces []NetworkInterface `gorm:"foreignkey:machine_id"`
}

func (m *Machine) LiveImagePaths(basepath string) (string, string) {
	return filepath.Join(basepath, m.ID+"-config.img"), filepath.Join(basepath, m.ID+".qcow2")
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
	BridgeID      uint32
}

func (n *Network) NetlinkVxlan() string {
	return fmt.Sprintf("vxlan-%d", n.BridgeID)
}

func (n *Network) NetlinkVxlanId() int {
	return int(n.BridgeID) % (1 << 24)
}

func (n *Network) NetlinkBridge() string {
	if n.BridgeID != 0 {
		return fmt.Sprintf("vxbr-%d", n.BridgeID)
	}
	return fmt.Sprintf("natbr-%s", n.Name)
}

func (n *Network) IsBridge() bool {
	return n.BridgeID != 0
}

type NetworkSpec struct {
	Subnet  string
	Gateway string
}
