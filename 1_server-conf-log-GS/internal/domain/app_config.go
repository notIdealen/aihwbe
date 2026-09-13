package domain

type AppConfig struct {
	Root     string `envconfig:""`
	LogDir   string
	LogLevel string
	Version  string
	HTTPAddr string
}

func NewAppConfig() *AppConfig {

	return &AppConfig{}
}
