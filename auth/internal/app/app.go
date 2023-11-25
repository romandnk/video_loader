package app

import (
	"log"

	"romandnk/video_loader/auth/config"
	zaplogger "romandnk/video_loader/auth/pkg/logger/zap"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	logger, err := zaplogger.NewLogger(cfg.ZapLogger)
	if err != nil {
		log.Fatalf("error initializing zap logger: %s", err.Error())
	}

	logger.Info("using zap logger")
}
