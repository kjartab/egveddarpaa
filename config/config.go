package config

import (
	"log"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpAddress string `envconfig:"HTTP_ADDRESS" default:"0.0.0.0:8080"`
}

var EnvPrefix = ""

func LoadEnvConfig() *Config {
	var cfg Config
	if err := envconfig.Process(EnvPrefix, &cfg); err != nil {
		log.Fatalf("config: Unable to load config for %T: %s", &cfg, err)
	}
	return &cfg
}
