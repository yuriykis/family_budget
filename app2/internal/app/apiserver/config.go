package apiserver

type Config struct {
	BindAddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
