package apiserver

type Config struct {
	BindAddr    string `yaml:"bind_addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8002",
		LogLevel: "debug",
	}
}
