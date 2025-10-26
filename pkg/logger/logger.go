package logger

import (
	"log/slog"
	"os"

	"github.com/cdxy1/go-file-storage/internal/config"
)

func SetupLogger() *slog.Logger {
	var log *slog.Logger
	cfg := config.GetConfig()

	switch cfg.Logger.Env {
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn}))
	case "local":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}
