package internal

import "github.com/caarlos0/env"

type Config struct {
	Port                  string `env:"PORT" envDefault:"8080"`
	HealthPort            string `env:"HEALTH_PORT" envDefault:"8090"`
	GqlMainEndpoint       string `env:"GQL_ENDPOINT" envDefault:"/"`
	GqlPlaygroundEndpoint string `env:"GQL_PLAYGROUND" envDefault:"/playground"`
	LogLevel              string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, env.Parse(cfg)
}
