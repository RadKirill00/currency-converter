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

	log := setupLogger(cfg.Env)

	log.Info("Starting", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

}



func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
		case envLocal:  
			log = slog.New(
				slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		case envDev: {
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
			)
		}
		case envProd: {
			log = slog.New(
				slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
			)
		}
	}
	return log
}
