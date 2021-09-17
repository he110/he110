package internal

import "github.com/caarlos0/env"

type Config struct {
	Port                  string `env:"PORT" envDefault:"8080"`
	GqlMainEndpoint       string `env:"GQL_ENDPOINT" envDefault:"/"`
	GqlPlaygroundEndpoint string `env:"GQL_PLAYGROUND" envDefault:"/playground"`
	HealthPort            string `env:"HEALTH_PORT" envDefault:"8090"`
	LogLevel              string `env:"LOG_LEVEL" envDefault:"debug"`
	DBUser                string `env:"DB_USER" envDefault:"user"`
	DBPassword            string `env:"DB_PASSWORD" envDefault:""`
	DBSchema              string `env:"DB_SCHEMA" envDefault:"personal"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, env.Parse(cfg)
}
