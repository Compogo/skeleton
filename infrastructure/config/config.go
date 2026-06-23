// Package config содержит конфигурацию самого скелетона.
//
// Этот пакет — пример того, как должна выглядеть конфигурация в любом сервисе на Compogo.
//
// Конфигурация состоит из:
//   - Структуры Config с полями
//   - Конструктора NewConfig
//   - Функции Configuration для загрузки из Configurator
//   - Компонента для регистрации в DI-контейнере
package config

import (
	"github.com/Compogo/compogo"
)

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

func Configuration(config *Config, configurator compogo.Configurator) *Config {
	if config.Test == "" || config.Test == TestDefault {
		configurator.SetDefault(TestFieldName, TestDefault)
		config.Test = configurator.GetString(TestFieldName)
	}

	return config
}
