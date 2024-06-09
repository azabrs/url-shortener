package main

import (
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDebug = "debug"
)


func main(){

	cfg := config.MustLoad("../config/config.yaml")
	log := SetupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
}

func SetupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env{
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDebug:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}