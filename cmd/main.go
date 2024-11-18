package main

import (
	"currency-converter/internal/config"
	"fmt"
	"log"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev = "dev"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

}



func setupLogger(env string) {
	var log *slog.Logger

	switch env {
	case envLocal:  
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
}
