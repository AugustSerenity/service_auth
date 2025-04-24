package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AugustSerenity/service_auth/internal/config"
	"github.com/AugustSerenity/service_auth/internal/handler"
	"github.com/AugustSerenity/service_auth/internal/service"
	"github.com/AugustSerenity/service_auth/internal/storage"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	configPath := flag.String("config", "config/config.yaml", "config file path")
	flag.Parse()

	cfg := config.ParseConfig(*configPath)

	db := storage.InitDB(cfg.DB)
	defer storage.CloseDB(db)

	storage := storage.New(db)

	srv := service.New(storage, []byte(cfg.Secret))

	h := handler.New(srv)

	s := http.Server{
		Addr:         cfg.Address,
		Handler:      h.Route(),
		IdleTimeout:  cfg.IdleTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Println("listen and serve error: %w", err)
		}
	}()

	log.Println("starting server: address", cfg.Server.Address)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	log.Println("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		log.Println("shutdown server error: %w", err)
	}
}
