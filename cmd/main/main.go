package main

import (
	"log"

	"github.com/m1kx/IHK_Abschluss_Projekt/internal/config"
	"github.com/m1kx/IHK_Abschluss_Projekt/internal/db"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New("config.yaml")
	if err != nil {
		return err
	}
	log.Printf("config loaded: %+v", cfg.Get())

	database, err := db.New(cfg.Get().Database.DSN)
	if err != nil {
		return err
	}
	defer database.Close()
	log.Print("database connected")
	return nil
}
