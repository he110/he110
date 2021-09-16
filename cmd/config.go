package main

import "github.com/caarlos0/env"

type Config struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, env.Parse(cfg)
}