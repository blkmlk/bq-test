package storage

import (
	"bq/internal/models"
	"bq/internal/services/config"
	"context"

	"github.com/sarulabs/di/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlStorage struct {
	config config.Config
	db     *gorm.DB
}

func newMysqlStorage(ctn di.Container) Storage {
	return &mysqlStorage{
		config: ctn.Get(config.DefinitionName).(config.Config),
	}
}

func (m *mysqlStorage) init() error {
	db, err := gorm.Open(mysql.Open(m.config.DBUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	m.db = db
	return nil
}

func (m *mysqlStorage) CreateRecord(ctx context.Context, record *models.Record) error {
	tx := m.db.WithContext(ctx).Create(record)
	return tx.Error
}

func (m *mysqlStorage) GetRecord(ctx context.Context, id int) (*models.Record, error) {
	var record models.Record
	tx := m.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&record)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &record, nil
}
