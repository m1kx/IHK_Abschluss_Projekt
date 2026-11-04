package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/m1kx/IHK_Abschluss_Projekt/internal/config"
	"github.com/m1kx/IHK_Abschluss_Projekt/internal/db"
	"github.com/m1kx/IHK_Abschluss_Projekt/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New("config.yaml")
	if err != nil {
		return err
	}

	database, err := db.New(cfg.Get().Database.DSN)
	if err != nil {
		return err
	}
	defer database.Close()

	for team := range cfg.Get().Application.Instances {
		metrics.InitTeam(team)
	}

	http.Handle("/metrics", promhttp.Handler())
	slog.Info("listening", "addr", cfg.Get().Server.Addr)
	return http.ListenAndServe(cfg.Get().Server.Addr, nil)
}
