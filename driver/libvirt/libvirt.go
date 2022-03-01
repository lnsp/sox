package libvirt

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"text/template"

	"github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	"github.com/lnsp/virtm/driver/cloudconfig"
	"github.com/lnsp/virtm/driver/models"
	"gopkg.in/yaml.v2"
)

type Libvirt struct {
	conn        *libvirt.Connect
	network     *libvirt.Network
	storagePool *libvirt.StoragePool
	storagePath string
}

func New(uri string, network string, storagePath string) (*Libvirt, error) {
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
	storagePool, err := conn.LookupStoragePoolByTargetPath(storagePath)
	if err != nil {
		return nil, fmt.Errorf("find storage pool: %w", err)
	}
	storagePoolId, err := storagePool.GetUUIDString()
	if err != nil {
		return nil, fmt.Errorf("get storage pool id: %w", err)
	}
	log.Println("found storage pool", storagePoolId)
	return &Libvirt{
		conn:        conn,
		network:     net,
		storagePool: storagePool,
		storagePath: storagePath,
	}, nil
}

var networkIfaceTemplate = template.Must(template.New("network").Parse(`
auto {{ .Name }}
iface {{ .Name }} inet static
	address {{ .AddressCIDR }}
	{{ if .Gateway }}gateway {{ .Gateway }}{{ end }}
	{{ if .Nameservers }}dns-nameservers{{ range .Nameservers }} {{ . }}{{ end }}{{ end }}
`))

func configureImageNetworkInterface(machine *models.Machine, image string) error {
	// Create single tempdir
	netdir, err := os.MkdirTemp("", machine.ID)
	if err != nil {
		return fmt.Errorf("create netcfg dir: %w", err)
	}
	// Create single tempfile
	netcfg, err := os.Create(filepath.Join(netdir, "10-netcfg"))
	if err != nil {
		return fmt.Errorf("create netcfg file: %w", err)
	}
	defer netcfg.Close()
	// Write configuration to netcfg
	for i, iface := range machine.NetworkInterfaces {
		if err := networkIfaceTemplate.Execute(netcfg, struct {
			Name        string
			AddressCIDR string
			Nameservers []string
			Gateway     string
		}{
			Name:        fmt.Sprintf("enp%ds0", i+1),
			Nameservers: strings.Fields(iface.Network.Nameservers),
			AddressCIDR: iface.IPv4,
			Gateway:     iface.Network.IPv4.Gateway,
		}); err != nil {
			return fmt.Errorf("write netcfg item: %w", err)
		}
	}
	// Use virt-customize to push config file into /etc/network/interfaces.d
	virtCustomizeCmd := exec.Command("virt-customize", "-a", image, "--copy-in", netcfg.Name()+":/etc/network/interfaces.d")
	virtCustomizeCmd.Stderr = log.Writer()
	if err := virtCustomizeCmd.Run(); err != nil {
		return fmt.Errorf("copy netcfg to vm: %w", err)
	}
	log.Println("configured image network", netcfg.Name())
	return nil
}

var usernamePattern = regexp.MustCompile(`^[a-z][-a-z0-9]*$`)

func writeCloudConfig(machine *models.Machine) (string, error) {
	// Generate authorized_keys file
	authorizedKeys := bytes.NewBuffer(nil)
	for _, key := range machine.SSHKeys {
		fmt.Fprintln(authorizedKeys, key.Pubkey)
	}
	// Make sure username is valid
	if !usernamePattern.MatchString(machine.User) {
		return "", fmt.Errorf("username must be valid")
	}
	// Create cloud config
	shortuuid := machine.ID[:8]
	cc := cloudconfig.CloudConfig{
		Hostname:       shortuuid,
		FQDN:           shortuuid,
		ManageEtcHosts: true,
		Users: []cloudconfig.User{
			{
				Name:           machine.User,
				Sudo:           "ALL=(ALL:ALL) NOPASSWD:ALL",
				Home:           "/home/" + machine.User,
				Shell:          "/bin/bash",
				LockPasswd:     false,
				AuthorizedKeys: authorizedKeys.String(),
			},
		},
		Chpasswd: cloudconfig.Chpasswd{
			List:   []string{"debian:debian"},
			Expire: false,
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

func buildDomXml(id string, specs models.Specs, configImage, osImage string, ifaces []models.NetworkInterface) string {
	// Generate network interface list
	lvIfaces := make([]libvirtxml.DomainInterface, len(ifaces))
	for i := range ifaces {
		source := &libvirtxml.DomainInterfaceSource{}
		if ifaces[i].Network.Bridge != "" {
			source.Bridge = &libvirtxml.DomainInterfaceSourceBridge{
				Bridge: ifaces[i].Network.Bridge,
			}
		} else {
			source.Network = &libvirtxml.DomainInterfaceSourceNetwork{
				Network: ifaces[i].Network.Name,
			}
		}
		lvIfaces[i] = libvirtxml.DomainInterface{
			Source: source,
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
	// Stop domain if necessary
	state, _, err := dom.GetState()
	if err != nil {
		return fmt.Errorf("get domain state: %w", err)
	}
	if state == libvirt.DOMAIN_RUNNING {
		// Stop domain
		if err := dom.Destroy(); err != nil {
			return fmt.Errorf("destroy domain: %w", err)
		}
		log.Println("destroyed libvirt domain", machine.ID)
	}
	// Undefine domain
	if err := dom.Undefine(); err != nil {
		return fmt.Errorf("undefine domain: %w", err)
	}
	log.Println("undefined libvirt domain", machine.ID)
	// Delete disks
	configDiskPath, imageDiskPath := machine.LiveImagePaths(lv.storagePath)
	if err := os.Remove(configDiskPath); err != nil {
		log.Println("attempted to delete config disk:", err)
	}
	if err := os.Remove(imageDiskPath); err != nil {
		log.Println("attempted to delete image disk:", err)
	}
	return nil
}

// RebootMachine reboots an active machine.
func (lv *Libvirt) RebootMachine(machine *models.Machine) error {
	dom, err := lv.conn.LookupDomainByUUIDString(machine.ID)
	if err != nil {
		return fmt.Errorf("lookup domain: %w", err)
	}
	if err := dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT); err != nil {
		return fmt.Errorf("reboot domain: %w", err)
	}
	return nil
}

// StopMachines stops an active machine.
func (lv *Libvirt) StopMachine(machine *models.Machine) error {
	dom, err := lv.conn.LookupDomainByUUIDString(machine.ID)
	if err != nil {
		return fmt.Errorf("lookup domain: %w", err)
	}
	if err := dom.Destroy(); err != nil {
		return fmt.Errorf("destroy domain: %w", err)
	}
	return nil
}

// StartMachine starts a new machine.
func (lv *Libvirt) StartMachine(machine *models.Machine) error {
	dom, err := lv.conn.LookupDomainByUUIDString(machine.ID)
	if err != nil {
		return fmt.Errorf("lookup domain: %w", err)
	}
	if err := dom.Create(); err != nil {
		return fmt.Errorf("create domain: %w", err)
	}
	return nil
}

// writeDisabledNetworkConfig creates a basic network config that disabled cloud-init networking setup.
func writeDisabledNetworkConfig() (string, error) {
	// Create basic network config
	netcfg, err := os.CreateTemp("", "netcfg")
	if err != nil {
		return "", fmt.Errorf("network config: %w", err)
	}
	defer netcfg.Close()
	fmt.Fprintln(netcfg, "network: { config: disabled }")
	return netcfg.Name(), nil
}

func (lv *Libvirt) GetMachineState(id string) (models.MachineState, error) {
	// No entry found, unlock and get entry
	dom, err := lv.conn.LookupDomainByUUIDString(id)
	if err != nil {
		return models.StateUnknown, fmt.Errorf("lookup domain: %w", err)
	}
	state, _, err := dom.GetState()
	if err != nil {
		return models.StateUnknown, fmt.Errorf("get domain state: %w", err)
	}
	var machineState models.MachineState
	switch state {
	case libvirt.DOMAIN_RUNNING:
		machineState = models.StateRunning
	case libvirt.DOMAIN_BLOCKED, libvirt.DOMAIN_CRASHED:
		machineState = models.StateCrashed
	case libvirt.DOMAIN_PAUSED, libvirt.DOMAIN_PMSUSPENDED, libvirt.DOMAIN_SHUTDOWN, libvirt.DOMAIN_SHUTOFF:
		machineState = models.StateStopped
	default:
		machineState = models.StateUnknown
	}
	return machineState, nil
}

func (lv *Libvirt) CreateMachine(machine *models.Machine) error {
	// Get source img path
	configImagePath, osImagePath := machine.LiveImagePaths(lv.storagePath)
	osImageSize := fmt.Sprintf("%dG", machine.Specs.Disk)
	// Create image snapshot
	if err := exec.Command("qemu-img", "create", "-b", machine.Image.Path, "-f", "qcow2", "-F", "qcow2", osImagePath, osImageSize).Run(); err != nil {
		return fmt.Errorf("create image snapshot: %w", err)
	}
	log.Println("replicated image", machine.Image.ID, "to", osImagePath)
	// Setup networking in snapshot
	if err := configureImageNetworkInterface(machine, osImagePath); err != nil {
		return fmt.Errorf("configure image network: %w", err)
	}
	netcfg, err := writeDisabledNetworkConfig()
	if err != nil {
		return fmt.Errorf("network config: %w", err)
	}
	// Create cloud config
	cloudcfg, err := writeCloudConfig(machine)
	if err != nil {
		return fmt.Errorf("cloud config: %w", err)
	}
	log.Println("created cloud config", cloudcfg)
	// Merge into image
	if err := exec.Command("cloud-localds", "-v", "-N", netcfg, configImagePath, cloudcfg).Run(); err != nil {
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
