package libvirt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
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

func writeCloudConfig(machine *models.Machine) (string, error) {
	var nameservers, searchDomains []string
	for _, iface := range machine.NetworkInterfaces {
		nameservers = append(nameservers, strings.Fields(iface.Network.Nameservers)...)
		searchDomains = append(searchDomains, strings.Fields(iface.Network.SearchDomains)...)
	}
	// Concatenate nameservers
	// Create cloud config
	shortuuid := machine.ID[:8]
	cc := cloudconfig.CloudConfig{
		Hostname:       shortuuid,
		FQDN:           shortuuid,
		ManageEtcHosts: true,

		Users: []cloudconfig.User{
			{
				Name:           "debian",
				Sudo:           "ALL=(ALL:ALL) NOPASSWD:ALL",
				Home:           "/home/debian",
				Shell:          "/bin/bash",
				LockPasswd:     false,
				AuthorizedKeys: machine.SSHKey.Pubkey,
			},
		},

		Chpasswd: cloudconfig.Chpasswd{
			List:   []string{"debian:debian"},
			Expire: false,
		},

		Packages:      []string{"qemu-guest-agent"},
		PackageUpdate: true,

		ManageResolvConf: true,
		ResolvConf: cloudconfig.ResolvConf{
			Nameservers:   nameservers,
			SearchDomains: searchDomains,
		},
	}
	content, err := yaml.Marshal(cc)
	if err != nil {
		return "", fmt.Errorf("create cloudconfig: %w", err)
	}
	ccTempFile, err := os.CreateTemp("", machine.ID)
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

func writeNetworkConfig(machine *models.Machine) (string, error) {
	ethernets := make(map[string]cloudconfig.NetworkEthernet)
	for i := range machine.NetworkInterfaces {
		ethernets[fmt.Sprintf("enp1s%d", i)] =
			cloudconfig.NetworkEthernet{
				DHCPv4:      false,
				Addresses:   []string{machine.NetworkInterfaces[i].IPv4},
				GatewayIPv4: machine.NetworkInterfaces[i].Network.IPv4.Gateway,
				Nameservers: cloudconfig.NetworkNameservers{
					Addresses:     strings.Fields(machine.NetworkInterfaces[i].Network.Nameservers),
					SearchDomains: strings.Fields(machine.NetworkInterfaces[i].Network.SearchDomains),
				},
			}
	}
	netcfg := cloudconfig.NetworkConfig{
		Version:   2,
		Ethernets: ethernets,
	}
	content, err := yaml.Marshal(netcfg)
	if err != nil {
		return "", fmt.Errorf("marshal network config: %w", err)
	}
	netcfgTempFile, err := os.CreateTemp("", machine.ID)
	if err != nil {
		return "", fmt.Errorf("create temp netconfig: %w", err)
	}
	defer netcfgTempFile.Close()
	if _, err := netcfgTempFile.Write(content); err != nil {
		return "", fmt.Errorf("write netconfig: %w", err)
	}
	return netcfgTempFile.Name(), nil
}

func buildDomXml(id string, specs models.Specs, configImage, osImage string, ifaces []models.NetworkInterface) string {
	// Generate network interface list
	lvIfaces := make([]libvirtxml.DomainInterface, len(ifaces))
	for i := range ifaces {
		lvIfaces[i] = libvirtxml.DomainInterface{
			Source: &libvirtxml.DomainInterfaceSource{
				Network: &libvirtxml.DomainInterfaceSourceNetwork{
					Network: ifaces[i].Network.Name,
				},
			},
			MAC: &libvirtxml.DomainInterfaceMAC{
				Address: ifaces[i].HwAddr,
			},
			Model: &libvirtxml.DomainInterfaceModel{
				Type: "virtio",
			},
		}
	}
	// Define domain
	dom := &libvirtxml.Domain{
		Type: "kvm",
		Name: id,
		UUID: id,
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Arch:    "x86_64",
				Machine: "q35",
				Type:    "hvm",
			},
			BootDevices: []libvirtxml.DomainBootDevice{
				{Dev: "hd"},
			},
			BootMenu: &libvirtxml.DomainBootMenu{
				Enable: "no",
			},
		},
		Features: &libvirtxml.DomainFeatureList{
			ACPI: &libvirtxml.DomainFeature{},
			APIC: &libvirtxml.DomainFeatureAPIC{},
		},
		CPU: &libvirtxml.DomainCPU{
			Mode: "host-model",
		},
		Clock: &libvirtxml.DomainClock{
			Offset: "utc",
			Timer: []libvirtxml.DomainTimer{
				{Name: "rtc", TickPolicy: "catchup"},
				{Name: "pit", TickPolicy: "delay"},
				{Name: "hpet", Present: "no"},
			},
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: uint(specs.CPUs),
		},
		Memory: &libvirtxml.DomainMemory{
			Value: uint(specs.Memory),
			Unit:  "M",
		},
		Devices: &libvirtxml.DomainDeviceList{
			Emulator: "/usr/bin/qemu-system-x86_64",
			Disks: []libvirtxml.DomainDisk{
				{
					Device: "cdrom",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "raw",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: configImage,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
						Bus: "sata",
					},
				},
				{
					Device: "disk",
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: osImage,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
				},
			},
			Interfaces: lvIfaces,
			Consoles: []libvirtxml.DomainConsole{
				{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "serial",
					},
				},
			},
			Graphics: []libvirtxml.DomainGraphic{
				{
					VNC: &libvirtxml.DomainGraphicVNC{
						Port: -1,
					},
				},
			},
			RNGs: []libvirtxml.DomainRNG{
				{
					Model: "virtio",
					Backend: &libvirtxml.DomainRNGBackend{
						Random: &libvirtxml.DomainRNGBackendRandom{
							Device: "/dev/urandom",
						},
					},
				},
			},
			Videos: []libvirtxml.DomainVideo{
				{
					Model: libvirtxml.DomainVideoModel{
						Type: "vga",
					},
				},
			},
		},
	}
	xml, _ := dom.Marshal()
	return xml
}

func (lv *Libvirt) DeleteMachine(machine *models.Machine) error {
	dom, err := lv.conn.LookupDomainByUUIDString(machine.ID)
	if err != nil {
		return fmt.Errorf("lookup domain: %w", err)
	}
	// Stop domain
	if err := dom.Destroy(); err != nil {
		return fmt.Errorf("destroy domain: %w", err)
	}
	log.Println("destroyed libvirt domain", machine.ID)
	// Undefine domain
	if err := dom.Undefine(); err != nil {
		return fmt.Errorf("undefine domain: %w", err)
	}
	log.Println("undefined libvirt domain", machine.ID)
	// TODO(lnsp): What about disks?
	return nil
}

func (lv *Libvirt) CreateMachine(machine *models.Machine) error {
	// Get source img path
	configImagePath := filepath.Join(filepath.Dir(machine.Image.Path), machine.ID+"-config.img")
	osImagePath := filepath.Join(filepath.Dir(machine.Image.Path), machine.ID+".qcow2")
	osImageSize := fmt.Sprintf("%dM", machine.Specs.Disk)
	// Create image snapshot
	if err := exec.Command("qemu-img", "create", "-b", machine.Image.Path, "-f", "qcow2", "-F", "qcow2", osImagePath, osImageSize).Run(); err != nil {
		return fmt.Errorf("create image snapshot: %w", err)
	}
	log.Println("replicated image", machine.Image.ID, "to", osImagePath)
	// Create cloud config
	cloudcfg, err := writeCloudConfig(machine)
	if err != nil {
		return fmt.Errorf("cloud config: %w", err)
	}
	log.Println("created cloud config", cloudcfg)
	// Create network config
	netcfg, err := writeNetworkConfig(machine)
	if err != nil {
		return fmt.Errorf("network config: %w", err)
	}
	log.Println("create network config", netcfg)
	// Merge into image
	if err := exec.Command("cloud-localds", "-v", "--network-config="+netcfg, configImagePath, cloudcfg).Run(); err != nil {
		return fmt.Errorf("merge config: %w", err)
	}
	// Generate domain xml
	domXml := buildDomXml(machine.ID, machine.Specs, configImagePath, osImagePath, machine.NetworkInterfaces)
	dom, err := lv.conn.DomainDefineXML(domXml)
	if err != nil {
		return fmt.Errorf("define domain: %w", err)
	}
	log.Println("defined libvirt domain", machine.ID)
	// And start domain
	if err := dom.Create(); err != nil {
		return fmt.Errorf("create domain: %w", err)
	}
	log.Println("created libvirt domain", machine.ID)
	return nil
}
