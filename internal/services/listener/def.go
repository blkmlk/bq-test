package listener

import (
	"context"

	"github.com/sarulabs/di/v2"
)

const DefinitionName = "listener"

var (
	Definition = di.Def{
		Name: DefinitionName,
		Build: func(ctn di.Container) (interface{}, error) {
			return newListener(ctn), nil
		},
	}
)

type Listener interface {
	Start(ctx context.Context) error
}
