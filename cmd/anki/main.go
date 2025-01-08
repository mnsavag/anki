package main

import (
	"context"
	"log"

	"github.com/mnsavag/anki.git/internal/anki"
	"github.com/mnsavag/anki.git/internal/anki/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to init config, %s", err.Error())
	}

	app := anki.NewApp(cfg)
	app.Start(context.Background())
}
