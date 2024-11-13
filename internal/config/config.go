package config

import "time"


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