package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Server   `yaml:"server"`
	SubsRepo `yaml:"subscriptions-repo"`
}

type Server struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     string `yaml:"port" env:"PORT"`
	LogLevel string `yaml:"log_level" env-default:"info"`
}

type SubsRepo struct {
	DriverName    string `yaml:"driver"`
	URI           string `yaml:"db_uri" env:"DB_URI"`
	MaxOpenConns  int    `yaml:"max_open_conns" env-default:"15"`
	MigrationsDir string `yaml:"migrations_dir" env:"MIGRATIONS_DIR"`
}

func New(path string) (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
