package httpserver

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"ADDR" required:"true"`
	// Level string `envconfig:"LEVEL" required:"true"`
}

func NewServerConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("HTTP", &config)
	if err != nil {
		fmt.Println("Newlogger error", err)
		return nil, err
	}
	return &config, nil
}
