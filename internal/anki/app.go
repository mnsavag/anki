package anki

import (
	"context"
	"fmt"

	"github.com/mnsavag/anki.git/internal/anki/api"
	"github.com/mnsavag/anki.git/internal/anki/config"
	appgrpc "github.com/mnsavag/anki.git/internal/anki/grpc"
	apphttp "github.com/mnsavag/anki.git/internal/anki/http"
	"github.com/mnsavag/anki.git/internal/anki/repository/sqlite"
	"github.com/mnsavag/anki.git/internal/anki/service"
	"github.com/mnsavag/anki.git/internal/lib/log"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type App struct {
	cfg        config.Config
	httpServer *apphttp.Server
	grpcServer *appgrpc.Server
	logger     *log.Logger
}

func NewApp(cfg config.Config, logger *log.Logger) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *App) Init(ctx context.Context) error {
	// storage
	dbConn, err := sqlite.NewSqliteConn(a.cfg.Database.DSN)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	// server service repo
	repo := sqlite.NewSqliteRepository(dbConn, *a.logger)
	ankiService := service.NewService(
		a.logger.WithField("component", "service"),
		repo,
	)

	// grpc server
	a.grpcServer = appgrpc.NewServer(a.cfg, a.logger)
	apiServer := api.NewServer(a.logger, ankiService)
	a.grpcServer.RegisterAnkiV1(apiServer)

	// http server
	a.httpServer = apphttp.NewServer(a.cfg, a.logger)
	err = a.httpServer.Register(ctx)
	if err != nil {
		return fmt.Errorf("httpServer.Register: %w", err)
	}

	return nil
}

func (a *App) Start(ctx context.Context) error {
	errG, _ := errgroup.WithContext(ctx)
	errG.Go(func() error {
		return a.grpcServer.Start()
	})
	errG.Go(func() error {
		return a.httpServer.Start()
	})
	return nil
}

func (a *App) Stop() error {
	a.grpcServer.Stop()
	if err := a.httpServer.Stop(); err != nil {
		return errors.Wrap(err, "cannot stop http servers")
	}

	return nil
}
