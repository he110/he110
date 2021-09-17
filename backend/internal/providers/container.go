package providers

import (
	"database/sql"
	"fmt"

	"He110/PersonalWebSite/internal"
	"He110/PersonalWebSite/internal/graph/resolvers"
	"He110/PersonalWebSite/internal/providers/gql_provider"
	"He110/PersonalWebSite/internal/providers/health_provider"
	"He110/PersonalWebSite/internal/providers/logger_provider"
	"He110/PersonalWebSite/internal/providers/manager/activity_manager"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	constructors := []interface{}{
		func() (*internal.Config, error) {
			return internal.NewConfig()
		},
		func (am *activity_manager.ActivityManager) *resolvers.Resolver {
			return resolvers.NewResolver(am)
		},
		func (db *sql.DB, l *zap.Logger) *activity_manager.ActivityManager {
			return activity_manager.NewActivityManager(db, l)
		},
		func(cfg *internal.Config, l *zap.Logger, r *resolvers.Resolver) *gql_provider.GqlServer {
			return gql_provider.NewGqlServer(r, cfg.Port, cfg.GqlMainEndpoint, cfg.GqlPlaygroundEndpoint, l)
		},
		func(cfg *internal.Config) (*zap.Logger, error) {
			return logger_provider.NewLogger(cfg.LogLevel)
		},
		func(cfg *internal.Config, l *zap.Logger) *health_provider.HealthServer {
			return health_provider.NewHealthServer(cfg.HealthPort, "/", l)
		},
		func(cfg *internal.Config) (*sql.DB, error) {
			conUrl := fmt.Sprintf("%s:%s@/%s", cfg.DBUser, cfg.DBPassword, cfg.DBSchema)
			return sql.Open("mysql", fmt.Sprintf(conUrl))
		},
	}

	for _, c := range constructors {
		if err := container.Provide(c); err != nil {
			return nil, err
		}
	}

	return container, nil
}
