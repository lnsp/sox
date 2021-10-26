package driver

import (
	"fmt"

	"github.com/valar/virtm/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Driver struct {
	api.UnimplementedVirtMServer

	db *gorm.DB
}

func initModels(db *gorm.DB) error {
	db.AutoMigrate(&Machine{}, &Image{}, &SSHKey{})
	return nil
}

func New(dsn string) (*Driver, error) {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := initModels(db); err != nil {
		return nil, fmt.Errorf("init models: %w", err)
	}
	return &Driver{
		db: db,
	}, nil
}
