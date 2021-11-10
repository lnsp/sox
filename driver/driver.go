package driver

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver/libvirt"
	"github.com/valar/virtm/driver/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Driver struct {
	api.UnimplementedVirtMServer

	db *gorm.DB
	lv *libvirt.Libvirt
}

func (driver *Driver) ListSSHKeys(ctx context.Context, request *api.ListSSHKeysRequest) (*api.ListSSHKeysResponse, error) {
	keys := []models.SSHKey{}
	if result := driver.db.Find(&keys); result.Error != nil {
		return nil, status.Errorf(codes.Internal, "retrieve ssh keys: %v", result.Error)
	}
	apiKeys := make([]*api.SSHKey, len(keys))
	for i := range keys {
		apiKeys[i] = &api.SSHKey{
			Id:     keys[i].ID,
			Name:   keys[i].Name,
			Pubkey: keys[i].Pubkey,
		}
	}
	return &api.ListSSHKeysResponse{
		Keys: apiKeys,
	}, nil
}

func (driver *Driver) ListImages(ctx context.Context, request *api.ListImagesRequest) (*api.ListImagesResponse, error) {
	images := []models.Image{}
	if result := driver.db.Find(&images); result.Error != nil {
		return nil, status.Errorf(codes.Internal, "retrieve images: %v", result.Error)
	}
	apiImages := make([]*api.Image, len(images))
	for i := range images {
		apiImages[i] = &api.Image{
			Id:     images[i].ID,
			Name:   images[i].Name,
			System: api.Image_OS(api.Image_OS_value[images[i].OS]),
		}
	}
	return &api.ListImagesResponse{
		Images: apiImages,
	}, nil
}

func (driver *Driver) ConfigureNetworkInterface(ctx context.Context, network models.Network) (models.NetworkInterface, error) {
	inc := func(ipp net.IP) net.IP {
		ip := make(net.IP, len(ipp))
		copy(ip, ipp)
		for j := len(ip) - 1; j >= 0; j-- {
			ip[j]++
			if ip[j] > 0 {
				break
			}
		}
		return ip
	}
	// Generate hw addr
	random := make([]byte, 3)
	if _, err := rand.Read(random); err != nil {
		return models.NetworkInterface{}, fmt.Errorf("generate hwaddr: %v", err)
	}
	hwAddr := net.HardwareAddr(append([]byte{0x52, 0x54, 0x00}, random...)).String()
	log.Println("generated hwaddr", hwAddr)
	// Generate ip addr
	var existingIfaces []models.NetworkInterface
	if err := driver.db.Where("network_id = ?", network.ID).Find(&existingIfaces).Error; err != nil {
		return models.NetworkInterface{}, fmt.Errorf("retrieve interfaces: %v", err)
	}
	// Find min mask value on ip addr
	ip, ipnet, err := net.ParseCIDR(network.IPv4.Subnet)
	if err != nil {
		return models.NetworkInterface{}, fmt.Errorf("parse subnet: %v", err)
	}
	// TODO(lnsp): Refine this algorithm
	blocked := make(map[string]struct{})
	// Block base ip and gateway from pool
	blocked[ip.String()] = struct{}{}
	blocked[network.IPv4.Gateway] = struct{}{}
	// Block every iface from pool
	for _, iface := range existingIfaces {
		ip, _, _ := net.ParseCIDR(iface.IPv4)
		blocked[ip.String()] = struct{}{}
	}
	log.Println("searching for ip in subnet", network.IPv4.Subnet, "and blocked IPs", blocked)
	// Find first IP that is contained in our ipnet and not blocked
	for ipnet.Contains(ip) {
		if _, ok := blocked[ip.String()]; !ok {
			break
		}
		ip = inc(ip)
	}
	if !ipnet.Contains(ip) {
		return models.NetworkInterface{}, fmt.Errorf("no capacity available: subnet %s is full", ipnet.String())
	}
	maskSize, _ := ipnet.Mask.Size()
	ipv4 := fmt.Sprintf("%s/%d", ip, maskSize)
	log.Println("found available IPv4 address", ipv4)
	return models.NetworkInterface{
		Network: network,
		HwAddr:  hwAddr,
		IPv4:    ipv4,
	}, nil
}

func (driver *Driver) CreateMachine(ctx context.Context, request *api.CreateMachineRequest) (*api.CreateMachineResponse, error) {
	// Retrieve SSH keys
	sshKeys := make([]models.SSHKey, len(request.SshKeyIds))
	if len(request.SshKeyIds) < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "must contain at least one ssh key ID")
	}
	for i := range request.SshKeyIds {
		if result := driver.db.Where("id = ?", request.SshKeyIds[i]).First(&sshKeys[i]); result.Error != nil {
			return nil, status.Errorf(codes.NotFound, "retrieve ssh key: %v", result.Error)
		}
	}
	// Retrieve image
	var image models.Image
	if result := driver.db.Where("id = ?", request.ImageId).First(&image); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "retrieve image: %v", result.Error)
	}
	// Retrieve network interfaces
	if len(request.NetworkIds) < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "must contain at least one network ID")
	}
	ifaces := make([]models.NetworkInterface, len(request.NetworkIds))
	for i := range request.NetworkIds {
		var network models.Network
		if err := driver.db.Where("id = ?", request.NetworkIds[i]).First(&network).Error; err != nil {
			return nil, status.Errorf(codes.NotFound, "retrieve network: %v", err)
		}
		iface, err := driver.ConfigureNetworkInterface(ctx, network)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "configure interface: %v", err)
		}
		ifaces[i] = iface
	}
	// Create entry for machine in DB
	specs := models.Specs{
		CPUs:   request.Specs.Cpus,
		Memory: request.Specs.Memory,
		Disk:   request.Specs.Disk,
	}
	machine := models.Machine{
		ID:                uuid.New().String(),
		Name:              request.Name,
		Image:             image,
		SSHKeys:           sshKeys,
		Specs:             specs,
		NetworkInterfaces: ifaces,
	}
	// Generate network interface
	if err := driver.db.Create(&machine).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "create machine record: %v", err)
	}
	log.Println("created machine record", machine.ID)
	if err := driver.lv.CreateMachine(&machine); err != nil {
		return nil, status.Errorf(codes.Internal, "create machine instance: %v", err)
	}
	log.Println("created machine instance", machine.ID)
	return &api.CreateMachineResponse{
		Id: machine.ID,
	}, nil
}

func (driver *Driver) GetMachineDetails(ctx context.Context, request *api.GetMachineDetailsRequest) (*api.GetMachineDetailsResponse, error) {
	// Use ID or name to find machine
	var machine models.Machine
	if err := driver.db.Preload("SSHKeys").Preload("NetworkInterfaces").Where("id = ? OR name = ?", request.Id, request.Id).First(&machine).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "retrieve machine: %v", err)
	}
	// TODO(lnsp): Retrieve status data from libvirt, maybe cache them
	apiNetworkInterfaces := make([]*api.NetworkInterface, len(machine.NetworkInterfaces))
	for i := range machine.NetworkInterfaces {
		apiNetworkInterfaces[i] = &api.NetworkInterface{
			NetworkId: machine.NetworkInterfaces[i].Network.ID,
			IpV4:      machine.NetworkInterfaces[i].IPv4,
			IpV6:      machine.NetworkInterfaces[i].IPv6,
		}
	}
	// Generate list of key ids
	sshKeyIds := make([]string, len(machine.SSHKeys))
	for i := range machine.SSHKeys {
		sshKeyIds[i] = machine.SSHKeys[i].ID
	}
	// Get state
	state, err := driver.lv.GetMachineState(machine.ID)
	if err != nil {
		state = models.StateUnknown
	}
	// Put info into protobuf
	return &api.GetMachineDetailsResponse{
		Machine: &api.Machine{
			Id:     machine.ID,
			Name:   machine.Name,
			Status: machineStateToApiStatus(state),
			Specs: &api.Machine_Specs{
				Cpus:   machine.Specs.CPUs,
				Memory: machine.Specs.Memory,
				Disk:   machine.Specs.Disk,
			},
			ImageId:   machine.Image.ID,
			SshKeyIds: sshKeyIds,
			Networks:  apiNetworkInterfaces,
		},
	}, nil
}

func machineStateToApiStatus(state models.MachineState) api.Machine_Status {
	switch state {
	case models.StateCreated:
		return api.Machine_CREATED
	case models.StateRunning:
		return api.Machine_RUNNING
	case models.StateCrashed:
		return api.Machine_CRASHED
	case models.StateStopped:
		return api.Machine_STOPPED
	default:
		return api.Machine_STATUS_UNSPECIFIED
	}
}

func (driver *Driver) ListMachines(ctx context.Context, request *api.ListMachinesRequest) (*api.ListMachinesResponse, error) {
	machines := []models.Machine{}
	if err := driver.db.Find(&machines).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "retrieve machines: %v", err)
	}
	apiMachines := make([]*api.Machine, len(machines))
	for i := range machines {
		// get state
		state, err := driver.lv.GetMachineState(machines[i].ID)
		if err != nil {
			log.Println("get machine state:", err)
			state = models.StateUnknown
		}
		apiMachines[i] = &api.Machine{
			Id:     machines[i].ID,
			Name:   machines[i].Name,
			Status: machineStateToApiStatus(state),
			Specs: &api.Machine_Specs{
				Cpus:   machines[i].Specs.CPUs,
				Memory: machines[i].Specs.Memory,
				Disk:   machines[i].Specs.Disk,
			},
			ImageId: machines[i].Image.ID,
		}
	}
	return &api.ListMachinesResponse{
		Machines: apiMachines,
	}, nil
}

func (driver *Driver) DeleteMachine(ctx context.Context, request *api.DeleteMachineRequest) (*api.DeleteMachineResponse, error) {
	// Destroy machine instance
	var machine models.Machine
	if err := driver.db.Where("id = ?", request.Id, request.Id).First(&machine).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "retrieve machine: %v", err)
	}
	if err := driver.lv.DeleteMachine(&machine); err != nil {
		return nil, status.Errorf(codes.Internal, "delete machine: %v", err)
	}
	// Delete machine record
	if err := driver.db.Select("NetworkInterfaces").Delete(&machine).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "delete machine record: %v", err)
	}
	return &api.DeleteMachineResponse{}, nil
}

func (driver *Driver) ListNetworks(ctx context.Context, request *api.ListNetworksRequest) (*api.ListNetworksResponse, error) {
	networks := []models.Network{}
	if err := driver.db.Find(&networks).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "retrieve networks: %v", err)
	}
	apiNetworks := make([]*api.Network, len(networks))
	for i := range networks {
		apiNetworks[i] = &api.Network{
			Id:   networks[i].ID,
			Name: networks[i].Name,
			IpV4: &api.Network_IpNetwork{
				Subnet:  networks[i].IPv4.Subnet,
				Gateway: networks[i].IPv4.Gateway,
			},
			IpV6: &api.Network_IpNetwork{
				Subnet:  networks[i].IPv6.Subnet,
				Gateway: networks[i].IPv6.Gateway,
			},
		}
	}
	return &api.ListNetworksResponse{
		Networks: apiNetworks,
	}, nil
}

func initModels(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.NetworkInterface{}, &models.Machine{}, &models.Image{}, &models.SSHKey{}, &models.Network{}); err != nil {
		return err
	}

	// make sure debian image exists
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Image{
		ID:   "6274bb3f-56c4-4a94-895b-8e0675f12368",
		Name: "Debian 11",
		OS:   api.Image_DEBIAN_BULLSEYE.String(),
		Path: "/var/lib/libvirt/images/debian-11-generic-amd64.qcow2",
	})
	// make sure ssh key exists
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.SSHKey{
		ID:     "f5e8f193-89b9-4557-b88d-f5dcb272577b",
		Name:   "default",
		Pubkey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPbrxB59qJKQ7WvJTEt9O8Esidzp6uuIhiPiMiyaHSrxf52R4H9CBqPBO/pC1AprBpk7ujI9YyjBaU7feig0w8xIRI04Uy9vknPTRDcKYIswjJqYu6i7CffjFnWz7Qj9/U/lSvOYV6qpicnx+jVX5aCnupMu8Qtt3udFN4Dnx5nW1hLwaIkBmzblNuGRZY3iYRKlSOijGavYGmNqTB809jBIr7B0+REI1C03zQbLGjQrXybBx0YZ3t+v7Cc/IG0kqBn94m3Q8oJ1yk7MWdMKYGB6iodPGKSfJ0TmlXdDIqPwL1LiHJCu3mRJzw/62iVrwxYYPjqnknzEQ6H2OhrvDtPAB6KqgIJ1V/exxwWYFglF4UUBkZZO8yiMIRQt+0E3NOTaV0uHawfyGsGvAZcphNCyYe5jBdRjolwEhaZCmre398ndL+e5CkjCnHMoAOLFFqCTIMseax/j04pyqcfiO4nP0+OssoEa1XrKWUMyGS6VHuFFbbthXN+/PQDA1x8n18Jnrql7AJrD71XqTYMwCoDY7Be/m4N8xIAqQPyt3/uP3XpkOeFvlJhJJM/uw7OeHtZraB7+CFbmpCKczhsz2xGV/YMiocxigrvEgUXZRSZKDvfLA4KDxDaxPdhDySvLRM0ZNcfSPkpYVdnIYco9x/p2NXyLN7TU/5D4K7GWutUQ== default",
	})
	// create default network
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Network{
		ID:   "eb7a6e41-da84-4db4-9cba-97509ddc8a58",
		Name: "network",
		IPv4: models.NetworkSpec{
			Subnet:  "192.168.100.0/24",
			Gateway: "192.168.100.1",
		},
		Nameservers: "1.1.1.1",
	})
	return nil
}

func New(dbDsn string, libvirtUri string) (*Driver, error) {
	db, err := gorm.Open(sqlite.Open(dbDsn))
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := initModels(db); err != nil {
		return nil, fmt.Errorf("init models: %w", err)
	}
	lv, err := libvirt.New(libvirtUri, "network", "default")
	if err != nil {
		return nil, fmt.Errorf("init libvirt: %w", err)
	}
	return &Driver{
		db: db,
		lv: lv,
	}, nil
}
