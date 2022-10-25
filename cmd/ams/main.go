package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/bots-empire/ams-service/internal/config"
	v1 "github.com/bots-empire/ams-service/internal/controller/http/v1"
	"github.com/bots-empire/ams-service/internal/db"
	"github.com/bots-empire/ams-service/internal/db/accesses"
	"github.com/bots-empire/ams-service/internal/httpserver"
	"github.com/bots-empire/ams-service/internal/log"
	"github.com/bots-empire/ams-service/internal/service"
)

func main() {
	// Init logger
	logger := log.NewProductionLogger(nil)
	printLogo(logger)

	// Init config
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Sugar().Fatalf("error init config: %v", err)
	}

	// Init database
	database, err := db.InitDataBase(context.Background(), cfg.RepositoryCfg)
	if err != nil {
		logger.Sugar().Fatalf("error init database: %v", err)
	}

	// Init Manager
	storage := accesses.NewStorage(database)
	manager := service.NewManager(logger, storage, []int64{1418862576}) // TODO: whitelist from config

	// Start HTTP server
	var mux = http.NewServeMux()
	v1.HandleRouts(mux, manager, logger)
	httpServer := httpserver.New(mux, httpserver.Port(cfg.ServicePort))
	logger.Sugar().Infof("server started on %s port", cfg.ServicePort)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Sugar().Warnf("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		logger.Sugar().Warnf("app - Run - httpServer.Notify: %v", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Sugar().Warnf("app - Run - httpServer.Shutdown: %v", err)
		return
	}
	logger.Sugar().Info("server shutdown")
}

func printLogo(logger *zap.Logger) {
	log.ClearTerminal()

	log.PrintLogo("AMS Service", []string{"DC71F5"})

	logger.Info("ams service is starting")
}
