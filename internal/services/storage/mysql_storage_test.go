package storage

import (
	"bq/internal/models"
	"bq/internal/services/config"
	"context"
	"testing"
	"time"

	"github.com/shopspring/decimal"

	"github.com/sarulabs/di/v2"

	"github.com/stretchr/testify/suite"
)

type MysqlStorageTestSuite struct {
	suite.Suite
	storage Storage
}

func TestMysqlStorage(t *testing.T) {
	suite.Run(t, new(MysqlStorageTestSuite))
}

func (t *MysqlStorageTestSuite) SetupSuite() {
	builder, err := di.NewBuilder()
	t.Require().NoError(err)

	t.Require().NoError(builder.Add(
		di.Def{
			Name: config.DefinitionName,
			Build: func(ctn di.Container) (interface{}, error) {
				return config.Config{
					DBUrl: "root:root@tcp(127.0.0.1:3306)/bq?charset=utf8mb4&parseTime=True&loc=Local",
				}, nil
			},
		},
		Definition,
	))

	ctn := builder.Build()
	t.storage = ctn.Get(DefinitionName).(Storage)
}

func (t *MysqlStorageTestSuite) TestCreateRecord() {
	ctx := context.Background()
	record := &models.Record{
		ID:           int(time.Now().Unix()),
		Change24Hour: decimal.RequireFromString("100"),
		CreatedAt:    time.Now(),
	}

	t.Require().NoError(t.storage.CreateRecord(ctx, record))

	foundRecord, err := t.storage.GetRecord(ctx, record.ID)
	t.Require().NoError(err)
	t.Require().Equal(record.ID, foundRecord.ID)
}
