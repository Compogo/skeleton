package cli

import (
	"context"

	"github.com/Compogo/compogo"
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/Compogo/compogo/logger"
	"github.com/Compogo/dig"
	"github.com/Compogo/logrus"
	"github.com/Compogo/sceleton/infrastructure/config"
	"github.com/Compogo/viper"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	app := compogo.NewApp(
		"skeleton",
		viper.WithViper(),
		dig.WithDig(),
		logrus.WithLogrus(),
		compogo.WithOsSignalCloser(),
		compogo.WithComponents(&component.Component{
			Name: "component",
			Dependencies: component.Components{
				config.Component,
			},
			Run: component.StepFunc(func(container container.Container) error {
				return container.Invoke(func(config *config.Config, logger logger.Logger) {
					logger.Infof("config test field value - %s", config.Test)
				})
			}),
			Wait: component.WaitFunc(func(ctx context.Context, container container.Container) error {
				return container.Invoke(func(logger logger.Logger) {
					logger.Info("press ctrl + C")

					<-ctx.Done()
				})
			}),
		}),
	)

	cmd := &cobra.Command{
		RunE: func(_ *cobra.Command, _ []string) error {
			return app.Serve()
		},
	}

	if err := app.BindFlags(cmd.PersistentFlags()); err != nil {
		return nil, err
	}

	return cmd, nil
}
