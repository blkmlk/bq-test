package publisher

import (
	"bq/internal/models"
	"context"

	"github.com/sarulabs/di/v2"
)

const DefinitionName = "publisher"

var (
	Definition = di.Def{
		Name: DefinitionName,
		Build: func(ctn di.Container) (interface{}, error) {
			p := newPublisher(ctn).(*publisher)

			if err := p.init(); err != nil {
				return nil, err
			}
			return p, nil
		},
	}
)

type Publisher interface {
	PublishRecord(ctx context.Context, record *models.Record) error
}
