package main

import (
	"context"
	"errors"
	"github.com/FrostyCreator/news-portal/internal/config"
	"github.com/FrostyCreator/news-portal/internal/server"
	"github.com/FrostyCreator/news-portal/internal/transport"
	"github.com/FrostyCreator/news-portal/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	configsDir = "configs/"
)

func main() {
	if err := logger.InitLogger(); err != nil {
		log.Panic(err)
	}

	cfg, err := config.GetConf(configsDir)
	if err != nil {
		logger.LogFatal(err.Error())
	}

	if cfg == nil {
		logger.LogFatal("Config is nil")
	}

	handlers := transport.NewHandler()

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.LogFatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.LogInfof("Server started with configs:\n%s\n ", cfg.String())

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.LogErrorf("failed to stop server: %v", err)
	}

}
