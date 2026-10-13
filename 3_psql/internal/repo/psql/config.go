package psql

import (
	"github.com/kelseyhightower/envconfig"
)

type PostgresConfig struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DBName   string `envconfig:"DB" required:"true"`
	Schema   string `envconfig:"SCHEMA" required:"true"`
	Table    string `envconfig:"TABLE" required:"true"`
	Timeout  int    `envconfig:"TIMER_FOR_CONTEXT" required:"true"`
}

func NewConfig() *PostgresConfig {
	var config PostgresConfig
	err := envconfig.Process("POSTGRES", &config)
	if err != nil {
		panic("Config for DB invalid")
	}
	return &config
}
