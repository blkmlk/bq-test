package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Host     string   `yaml:"host"`
	Interval int      `yaml:"interval"`
	URL      string   `yaml:"url"`
	DBUrl    string   `yaml:"db_url"`
	FSyms    []string `yaml:"fsyms"`
	TSyms    []string `yaml:"tsyms"`
}

func loadConfig(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	if err = yaml.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
