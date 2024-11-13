package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env string `yaml: "env"`
	StoragePath string `yaml: "storage" env-required: "true" `
	HTTPServer `yaml: "http_server"`
}

type HTTPServer struct {
	Adress string `yaml: "address" env-default: "localhost:8080"`
	Timeout time.Duration `yaml: "timeout"`
	IdleTimeout time.Duration `yaml: "idle-timeout"`
}

func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
}