package anki

import (
	"context"
	"fmt"

	"github.com/mnsavag/anki.git/internal/anki/config"
	apphttp "github.com/mnsavag/anki.git/internal/anki/http"
	"github.com/mnsavag/anki.git/internal/anki/repository/sqlite"
	"github.com/mnsavag/anki.git/internal/lib/log"
)

type App struct {
	cfg        config.Config
	httpServer *apphttp.Server
	logger     *log.Logger
}

func NewApp(cfg config.Config, logger *log.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	// storage
	_, err := sqlite.NewSqliteConn(a.cfg.Database.DSN)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	// server
	a.httpServer = apphttp.NewServer(a.cfg)
	return a.httpServer.Start()
}
