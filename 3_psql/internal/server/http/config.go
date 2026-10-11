package server

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Addr string `envconfig:"ADDR" required:"true"`
	//прочие таймауты
}

func NewConfigMust() *Config {
	var config Config
	if err := envconfig.Process("HOST", &config); err != nil {
		panic("Config for server invalid, check .env")
	}
	return &config
}
