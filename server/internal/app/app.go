// Package app configures and runs application.
package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nasdda/linkstore/config"
	v1 "github.com/nasdda/linkstore/internal/controller/http/v1"
	"github.com/nasdda/linkstore/internal/repo"
	"github.com/nasdda/linkstore/internal/usecase"
	"github.com/nasdda/linkstore/pkg/httpserver"
	"github.com/nasdda/linkstore/pkg/logger"
	"github.com/nasdda/linkstore/pkg/mongo"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.NewLogger("[APP]")

	// Dependencies
	mongoClient, err := mongo.NewClient(cfg.MongoConnectionString, context.Background())
	if err != nil {
		l.Errorf("failed to create mongo client: %w", err)
		return
	}
	mdb := mongoClient.Database(cfg.MongoDBName)

	// Repository
	linkRepo := repo.NewLinkRepo(mdb)

	// Use case
	link := usecase.NewLinkUseCase(linkRepo)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, link)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Errorf("app - Run - httpServer.Notify: %w", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Errorf("app - Run - httpServer.Shutdown: %w", err)
	}
}
