package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Init() (*Config, error) {

	dir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(dir)
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	cfg := Config{}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
