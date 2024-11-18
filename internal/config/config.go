package config

import (
	
	"time"

    "github.com/kelseyhightower/envconfig"
)


type Config struct {
	Env         string        `envconfig:"ENV" default:"local"`
	StoragePath string        `envconfig:"STORAGE_PATH" default:"./storage/storage.db"`
	Address     string        `envconfig:"ADDRESS" default:"localhost:8080"`
	Timeout     time.Duration `envconfig:"TIMEOUT" default:"4s"`
	IdleTimeout time.Duration `envconfig:"IDLE_TIMEOUT" default:"60s"`
}


func GetConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return Config{}, err
	}
	return config, nil
}