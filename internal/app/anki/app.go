package anki

import (
	"context"
	"fmt"
	"log"

	"github.com/mnsavag/anki.git/internal/anki/config"
	apphttp "github.com/mnsavag/anki.git/internal/anki/http"
	"github.com/mnsavag/anki.git/internal/anki/repository/sqlite"
)

type App struct {
	cfg        config.Config
	httpServer *apphttp.Server
}

func NewApp(cfg config.Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (a *App) Start(ctx context.Context) error {
	// storage
	_, err := sqlite.NewSqliteConn(a.cfg.Database.DSN)
	if err != nil {
		log.Fatalf("failed to init storage, %s", err.Error())
	}

	// server
	a.httpServer = apphttp.NewServer(a.cfg)
	if err := a.httpServer.Start(); err != nil {
		fmt.Println("Server start error")
	}
	return nil
}
