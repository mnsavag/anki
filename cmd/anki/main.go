package main

import (
	"context"
	logBase "log"
	"os"
	"os/signal"
	"syscall"

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
	if err := app.Init(context.Background()); err != nil {
		// logger.WithError(err).Fatal("init app failed")
		logBase.Fatalln(err)
	}

	if err := app.Start(context.Background()); err != nil {
		logBase.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	if err := app.Stop(); err != nil {
		logBase.Fatalln(err)
	}
}
