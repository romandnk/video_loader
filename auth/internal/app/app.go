package app

import (
	"context"
	"log"
	"os/signal"
	"romandnk/video_loader/auth/config"
	"syscall"

	"go.uber.org/zap"

	zaplogger "romandnk/video_loader/auth/pkg/logger/zap"
	"romandnk/video_loader/auth/pkg/storage/postgres"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	defer cancel()

	// initializing config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	// initializing zap logger
	logger, err := zaplogger.NewLogger(cfg.ZapLogger)
	if err != nil {
		log.Fatalf("error initializing zap logger: %s", err.Error())
	}

	logger.Info("using zap logger")

	// initializing connection to postgres db
	db, err := postgres.NewStorage(ctx, cfg.Postgres)
	if err != nil {
		db.Close()
		logger.Fatal("error initializing postgres db", zap.Error(err))
	}
	defer db.Close()

	logger.Info("using postgres storage",
		zap.String("host", cfg.Postgres.Host),
		zap.Int("port", cfg.Postgres.Port),
	)
}
