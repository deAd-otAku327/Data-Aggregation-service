package config

type Config struct {
	Server     `yaml:"server"`
	PostgresDB `yaml:"postgres-db"`
}

type Server struct {
	Host     string `yaml:"host" env:"HOST"`
	Port     string `yaml:"port" env:"PORT"`
	LogLevel string `yaml:"log_level" env-default:"info"`
}

type PostgresDB struct {
	DriverName    string `yaml:"driver"`
	URI           string `yaml:"db_uri" env:"DB_URI"`
	MaxOpenConns  int    `yaml:"max_open_conns" env-default:"15"`
	MigrationsDir string `yaml:"migrations_dir" env:"MIGRATIONS_DIR"`
}
