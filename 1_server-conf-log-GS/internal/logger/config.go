package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Dir   string `envconfig:"DIR" required:"true"`
	Level string `envconfig:"LEVEL" required:"true"`
}

func NewLoggerConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("LOG", &config)
	if err != nil {
		fmt.Println("Newlogger error", err)
		return nil, err
	}
	return &config, nil
}
