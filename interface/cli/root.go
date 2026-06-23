// Package cli предоставляет интерфейс командной строки для скелетона.
//
// Использует Cobra для создания CLI-команд.
// Вся инфраструктура (логгер, конфиг, DI, graceful shutdown) поднимается внутри команды.
package cli

import (
	"context"
	"time"

	"github.com/Compogo/compogo"
	"github.com/Compogo/dig"
	"github.com/Compogo/logrus"
	"github.com/Compogo/repeater"
	"github.com/Compogo/sceleton/infrastructure/config"
	"github.com/Compogo/viper"
	"github.com/spf13/cobra"
)

// NewRootCommand создаёт корневую команду CLI.
//
// В этой функции:
//  1. Создаётся приложение Compogo.
//  2. Подключаются стандартные компоненты (логгер, конфиг, DI, Closer).
//  3. Добавляются компоненты скелетона.
//  4. Создаётся Cobra-команда с RunE, запускающим app.Serve().
//  5. Привязываются флаги к приложению.
//
// Важно: команда создаётся, но не выполняется. Выполнение происходит в main.
func NewRootCommand() (*cobra.Command, error) {
	// Создаём приложение Compogo.
	// Это ядро, в которое мы добавляем все компоненты.
	app := compogo.NewApp(
		"skeleton", // имя приложения (используется в логах)

		// Стандартные компоненты
		viper.WithViper(),            // конфигурация через Viper
		dig.WithDig(),                // DI-контейнер через Dig
		logrus.WithLogrus(),          // логирование через Logrus
		compogo.WithOsSignalCloser(), // graceful shutdown через сигналы ОС

		// Компоненты скелетона
		compogo.WithComponents(
			// Компонент с примером использования репета
			&compogo.Component{
				Name: "example_component",
				Dependencies: compogo.Components{
					&config.Component,   // зависит от конфигурации
					&repeater.Component, // зависит от репета
				},
				// PreExecute — выполняется перед Execute.
				// Здесь можно подготовить данные, проверить зависимости.
				PreExecute: compogo.StepFunc(func(container compogo.Container) error {
					return container.Invoke(func(config *config.Config, logger compogo.Logger) {
						logger.Infof("config test field value - %s", config.Test)
					})
				}),
				// Execute — основной рабочий цикл.
				// Здесь запускаются фоновые задачи, серверы, воркеры.
				Execute: compogo.StepFunc(func(container compogo.Container) error {
					return container.Invoke(func(r repeater.Repeater, logger compogo.Logger, config *config.Config) error {
						// Добавляем задачу в Repeater (периодическое выполнение).
						return r.AddProcess(repeater.NewTask("test", time.Second, func(ctx context.Context) error {
							logger.Infof("log from task: config test field value - %s", config.Test)
							return nil
						}))
					})
				}),
				// Wait — ожидание сигнала завершения.
				// Здесь приложение блокируется до получения сигнала (Ctrl+C).
				Wait: compogo.WaitFunc(func(ctx context.Context, container compogo.Container) error {
					return container.Invoke(func(logger compogo.Logger) {
						logger.Info("press ctrl + C")
						<-ctx.Done()
					})
				}),
			},
		),
	)

	// Создаём корневую команду Cobra.
	cmd := &cobra.Command{
		Use:   "skeleton",
		Short: "Skeleton service for Compogo framework",
		Long:  "Example service demonstrating Compogo framework capabilities",
		RunE: func(_ *cobra.Command, _ []string) error {
			// Запускаем приложение.
			return app.Serve()
		},
	}

	// Привязываем флаги приложения к флагам команды.
	// Это позволяет использовать флаги через Cobra.
	if err := app.BindFlags(cmd.PersistentFlags()); err != nil {
		return nil, err
	}

	return cmd, nil
}
