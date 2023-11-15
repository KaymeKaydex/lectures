package config

type Config struct {
	Port int
}

func New() (*Config, error) {
	return &Config{
		Port: 8080,
	}, nil
}
