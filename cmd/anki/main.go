package main

import (
	"context"
	logBase "log"

	"github.com/mnsavag/anki.git/internal/anki"
	"github.com/mnsavag/anki.git/internal/anki/config"
	"github.com/mnsavag/anki.git/internal/lib/log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logBase.Fatalln(err)
	}

	logger, err := log.NewLogger(&cfg.Logger)
	if err != nil {
		logBase.Fatalln(err)
	}

	app := anki.NewApp(cfg, logger)
	app.Start(context.Background())
}
