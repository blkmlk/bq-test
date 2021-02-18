package config

import "github.com/sarulabs/di/v2"

const (
	DefinitionName           = "config"
	DefinitionNameConfigPath = "config-path"
)

var (
	Definition = di.Def{
		Name: DefinitionName,
		Build: func(ctn di.Container) (interface{}, error) {
			configPath := ctn.Get(DefinitionNameConfigPath).(string)
			return loadConfig(configPath)
		},
	}
)
