package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   LoggerConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type LoggerConfig struct {
	Env string
}

func GetConfig() *Config {
	v, err := LoadConfig("config-dev", "yaml")

	if err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}

	cfg, err := ParseConfig(v)

	if err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadConfig(filename, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.SetConfigType(filetype)
	v.AddConfigPath("./config")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
