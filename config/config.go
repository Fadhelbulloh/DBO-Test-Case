package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost        string `envconfig:"DB_HOST" json:"db_host"`
	DBPort        int    `envconfig:"DB_PORT" json:"db_port"`
	DBUser        string `envconfig:"DB_USER" json:"db_user"`
	DBPassword    string `envconfig:"DB_PASSWORD" json:"db_password"`
	DBName        string `envconfig:"DB_NAME" json:"db_name"`
	JWTSecret     string `envconfig:"JWT_SECRET" json:"jwt_secret"`
	JWTExpiration int    `envconfig:"JWT_EXPIRATION" json:"jwt_expiration"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("dev", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
