package config

import (
	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/flag"
)

// Component — компонент для регистрации конфигурации в DI-контейнере.
//
// Жизненный цикл компонента:
//  1. Init — регистрирует конструктор NewConfig в DI-контейнере.
//  2. BindFlags — привязывает флаг командной строки.
//  3. Configuration — загружает конфигурацию из Configurator.
//
// Это шаблон, который повторяется для каждой конфигурации в Compogo.
var Component = compogo.Component{
	Name: "skeleton.Config",
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(NewConfig)
	}),
	BindFlags: compogo.BindFlags(func(flagSet flag.FlagSet, container compogo.Container) error {
		return container.Invoke(func(config *Config) {
			flagSet.StringVar(&config.Test, TestFieldName, TestDefault, "")
		})
	}),
	Configuration: compogo.StepFunc(func(container compogo.Container) error {
		return container.Invoke(Configuration)
	}),
}
