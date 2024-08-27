package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type DatabaseConfig struct {
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"5432"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"name"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig("config/config.yaml", &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config, nil
}
