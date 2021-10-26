package driver

import (
	"context"
	"fmt"

	"github.com/valar/virtm/api"
	"github.com/valar/virtm/driver/libvirt"
	"github.com/valar/virtm/driver/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
		return nil, grpc.Errorf(codes.Internal, "retrieve ssh keys: %v", result.Error)
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
		return nil, grpc.Errorf(codes.Internal, "retrieve images: %v", result.Error)
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

func (driver *Driver) CreateMachine(ctx context.Context, request *api.CreateMachineRequest) (*api.CreateMachineResponse, error) {
	var sshKey models.SSHKey
	if result := driver.db.Where("id = ?", request.SshKeyId).First(&sshKey); result.Error != nil {
		return nil, grpc.Errorf(codes.NotFound, "retrieve ssh key: %v", result.Error)
	}
	var image models.Image
	if result := driver.db.Where("id = ?", request.ImageId).First(&image); result.Error != nil {
		return nil, grpc.Errorf(codes.NotFound, "retrieve image: %v", result.Error)
	}
	// create entry in database
	machine := models.Machine{}

	return nil, nil
}

func initModels(db *gorm.DB) error {
	db.AutoMigrate(&models.Machine{}, &models.Image{}, &models.SSHKey{})

	// make sure debian image exists
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Image{
		ID:   "6274bb3f-56c4-4a94-895b-8e0675f12368",
		Name: "debian/bullseye",
		OS:   api.Image_DEBIAN_BULLSEYE.String(),
		Path: "/var/lib/libvirt/images/debian-11-generic-amd64.qcow2",
	})
	// make sure ssh key exists
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.SSHKey{
		ID:     "f5e8f193-89b9-4557-b88d-f5dcb272577b",
		Name:   "default",
		Pubkey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPbrxB59qJKQ7WvJTEt9O8Esidzp6uuIhiPiMiyaHSrxf52R4H9CBqPBO/pC1AprBpk7ujI9YyjBaU7feig0w8xIRI04Uy9vknPTRDcKYIswjJqYu6i7CffjFnWz7Qj9/U/lSvOYV6qpicnx+jVX5aCnupMu8Qtt3udFN4Dnx5nW1hLwaIkBmzblNuGRZY3iYRKlSOijGavYGmNqTB809jBIr7B0+REI1C03zQbLGjQrXybBx0YZ3t+v7Cc/IG0kqBn94m3Q8oJ1yk7MWdMKYGB6iodPGKSfJ0TmlXdDIqPwL1LiHJCu3mRJzw/62iVrwxYYPjqnknzEQ6H2OhrvDtPAB6KqgIJ1V/exxwWYFglF4UUBkZZO8yiMIRQt+0E3NOTaV0uHawfyGsGvAZcphNCyYe5jBdRjolwEhaZCmre398ndL+e5CkjCnHMoAOLFFqCTIMseax/j04pyqcfiO4nP0+OssoEa1XrKWUMyGS6VHuFFbbthXN+/PQDA1x8n18Jnrql7AJrD71XqTYMwCoDY7Be/m4N8xIAqQPyt3/uP3XpkOeFvlJhJJM/uw7OeHtZraB7+CFbmpCKczhsz2xGV/YMiocxigrvEgUXZRSZKDvfLA4KDxDaxPdhDySvLRM0ZNcfSPkpYVdnIYco9x/p2NXyLN7TU/5D4K7GWutUQ== default",
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
