package logger

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Level string `envconfig:"LEVEL" default:"INFO"`
	Dir   string `envconfig:"DIR" required:"true"`
}

// config contains level and dir
func NewConfigMust() *Config {
	var config Config
	if err := envconfig.Process("LOGGER", &config); err != nil {
		panic("Config for logger invalid")
	}
	return &config
}
