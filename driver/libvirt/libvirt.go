package libvirt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/libvirt/libvirt-go"
	"github.com/valar/virtm/driver/cloudconfig"
	"github.com/valar/virtm/driver/models"
	"gopkg.in/yaml.v2"
)

type Libvirt struct {
	conn    *libvirt.Connect
	network *libvirt.Network
	storage *libvirt.StoragePool
}

func New(uri string, network string, storagePool string) (*Libvirt, error) {
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return nil, fmt.Errorf("connect to libvirt: %w", err)
	}
	log.Println("connected to libvirt", uri)
	// get network info
	net, err := conn.LookupNetworkByName(network)
	if err != nil {
		return nil, fmt.Errorf("find network: %w", err)
	}
	netId, _ := net.GetUUIDString()
	log.Println("found network", netId)
	// get storagepool info
	storage, err := conn.LookupStoragePoolByName(storagePool)
	if err != nil {
		return nil, fmt.Errorf("find storage pool: %w", err)
	}
	storageId, _ := storage.GetUUIDString()
	log.Println("found storage pool", storageId)
	return &Libvirt{conn, net, storage}, nil
}

func writeCloudConfig(id string, pubKey string) (string, error) {
	// Create cloud config
	cc := cloudconfig.CloudConfig{
		SetHostnameModule: cloudconfig.SetHostnameModule{
			Hostname: "debian",
			FQDN:     "debian.network",
		},
		ManageEtcHostsModule: cloudconfig.ManageEtcHostsModule{
			ManageEtcHosts: true,
		},
		UsersModule: cloudconfig.UsersModule{
			Users: []cloudconfig.User{
				{
					Name:           "debian",
					Sudo:           "(ALL:ALL) NOPASSWD:ALL",
					Home:           "/home/debian",
					Shell:          "/bin/bash",
					LockPasswd:     false,
					AuthorizedKeys: pubKey,
				},
			},
		},
		PackageModule: cloudconfig.PackageModule{
			Packages:      []string{"qemu-guest-agent"},
			PackageUpdate: true,
		},
	}
	content, err := yaml.Marshal(cc)
	if err != nil {
		return "", fmt.Errorf("create cloudconfig: %w", err)
	}
	ccTempFile, err := os.CreateTemp("", id)
	if err != nil {
		return "", fmt.Errorf("create temp cloudconfig: %w", err)
	}
	defer ccTempFile.Close()
	fmt.Fprintln(ccTempFile, "#cloud-config")
	if _, err := ccTempFile.Write(content); err != nil {
		return "", fmt.Errorf("write cloudconfig: %w", err)
	}
	return ccTempFile.Name(), nil
}

func writeNetworkConfig(id string, ipAddresses []string, gateway string, nameserver []string) (string, error) {
	netcfg := cloudconfig.NetworkConfig{
		Version: 2,
		Ethernets: map[string]cloudconfig.NetworkEthernet{
			"eth0": {
				DHCPv4:      false,
				Addresses:   ipAddresses,
				GatewayIPv4: gateway,
				Nameservers: cloudconfig.NetworkNameservers{
					Addresses: nameserver,
				},
			},
		},
	}
	content, err := yaml.Marshal(netcfg)
	if err != nil {
		return "", fmt.Errorf("marshal network config: %w", err)
	}
	netcfgTempFile, err := os.CreateTemp("", id)
	if err != nil {
		return "", fmt.Errorf("create temp netconfig: %w", err)
	}
	defer netcfgTempFile.Close()
	if _, err := netcfgTempFile.Write(content); err != nil {
		return "", fmt.Errorf("write netconfig: %w", err)
	}
	return netcfgTempFile.Name(), nil
}

func (lv *Libvirt) CreateMachine(id string, sshKey models.SSHKey, image models.Image, specs models.Specs) error {
	// Get source img path
	ccImagePath := filepath.Join(filepath.Dir(image.Path), id+"-config.img")
	destImagePath := filepath.Join(filepath.Dir(image.Path), id+".qcow2")
	destImageSize := fmt.Sprintf("%dM", specs.Disk)
	// Create image snapshot
	if err := exec.Command("qemu-img", "create", "-b", image.Path, "-f", "qcow2", "-F", "qcow2", destImagePath, destImageSize).Run(); err != nil {
		return fmt.Errorf("create image snapshot: %w", err)
	}
	// Create cloud config
	cloudcfg, err := writeCloudConfig(id, sshKey.Pubkey)
	if err != nil {
		return fmt.Errorf("cloud config: %w", err)
	}
	log.Println("created cloud config", cloudcfg)
	// Create network config
	netcfg, err := writeNetworkConfig(id, []string{"192.168.100.22/24"}, "192.168.100.1", []string{"192.168.100.1", "1.1.1.1"})
	if err != nil {
		return fmt.Errorf("network config: %w", err)
	}
	log.Println("create network config", netcfg)
	// Merge into image
	if err := exec.Command("cloud-localds", "-v", "--network-config="+netcfg, ccImagePath, cloudcfg).Run(); err != nil {
		return fmt.Errorf("merge config: %w", err)
	}
	return nil
}
