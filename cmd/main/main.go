package main

import (
	"log"

	"github.com/m1kx/IHK_Abschluss_Projekt/internal/config"
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
	return nil
}
