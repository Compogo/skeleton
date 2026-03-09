package config

import "github.com/Compogo/compogo/configurator"

const (
	TestFieldName = "app.test"

	TestDefault = "test"
)

type Config struct {
	Test string
}

func NewConfig() *Config {
	return &Config{}
}

func Configuration(config *Config, configurator configurator.Configurator) *Config {
	if config.Test == "" || config.Test == TestDefault {
		configurator.SetDefault(TestFieldName, TestDefault)
		config.Test = configurator.GetString(TestFieldName)
	}

	return config
}
