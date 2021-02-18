package listener

import (
	"bq/internal/models"
	"bq/internal/services/config"
	"bq/internal/services/publisher"
	"bq/internal/services/storage"
	"context"
	"time"

	"github.com/sarulabs/di/v2"
)

const (
	defaultInterval = 5
)

type listener struct {
	config    config.Config
	storage   storage.Storage
	publisher publisher.Publisher
}

func newListener(ctn di.Container) Listener {
	return &listener{
		config:    ctn.Get(config.DefinitionName).(config.Config),
		storage:   ctn.Get(storage.DefinitionName).(storage.Storage),
		publisher: ctn.Get(publisher.DefinitionName).(publisher.Publisher),
	}
}

func (l *listener) Start(ctx context.Context) error {
	interval := l.config.Interval

	if interval <= 0 {
		interval = defaultInterval
	}

	ticker := time.NewTicker(time.Second * time.Duration(interval))
	for {
		select {
		case <-ticker.C:
			err := l.loadRecords(ctx)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (l *listener) loadRecords(ctx context.Context) error {
	records, err := GetRecords(ctx, l.config.URL, l.config.FSyms, l.config.TSyms)
	if err != nil {
		return err
	}

	for fSymb, tRec := range records {
		for tSymb, record := range tRec {
			r := &models.Record{
				FSymb:           fSymb,
				TSymb:           tSymb,
				Change24Hour:    record.Change24Hour,
				ChangePCT24Hour: record.ChangePCT24Hour,
				Open24Hour:      record.Open24Hour,
				Volume24Hour:    record.Volume24Hour,
				Low24Hour:       record.Low24Hour,
				High24Hour:      record.High24Hour,
				Price:           record.Price,
				Supply:          record.Supply,
				MKTCAP:          record.MKTCAP,
				CreatedAt:       time.Now(),
			}

			err = l.storage.CreateRecord(ctx, r)
			if err != nil {
				return err
			}

			err = l.publisher.PublishRecord(ctx, r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
