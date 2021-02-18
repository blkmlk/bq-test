package storage

import (
	"bq/internal/models"
	"context"

	"github.com/sarulabs/di/v2"
)

const DefinitionName = "storage"

var (
	Definition = di.Def{
		Name: DefinitionName,
		Build: func(ctn di.Container) (interface{}, error) {
			storage := newMysqlStorage(ctn).(*mysqlStorage)

			if err := storage.init(); err != nil {
				return nil, err
			}

			return storage, nil
		},
	}
)

type Storage interface {
	CreateRecord(ctx context.Context, record *models.Record) error
	GetRecord(ctx context.Context, id int) (*models.Record, error)
}
