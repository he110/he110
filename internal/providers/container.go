package providers

import (
	"He110/PersonalWebSite/internal"
	"He110/PersonalWebSite/internal/providers/gql_provider"
	"He110/PersonalWebSite/internal/providers/logger_provider"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	constructors := []interface{}{
		func () (*internal.Config, error) {
			return internal.NewConfig()
		},
		func (cfg *internal.Config, l *zap.Logger) *gql_provider.GqlServer {
			return gql_provider.NewGqlServer(cfg.Port, cfg.GqlMainEndpoint, cfg.GqlPlaygroundEndpoint, l)
		},
		func (cfg *internal.Config) (*zap.Logger, error) {
			return logger_provider.NewLogger(cfg.LogLevel)
		},
	}

	for _, c := range constructors {
		if err := container.Provide(c); err != nil {
			return nil, err
		}
	}

	return container, nil
}
