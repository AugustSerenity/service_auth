package main

import (
	"flag"
	"net/http"

	"github.com/AugustSerenity/service_auth/internal/config"
	"github.com/AugustSerenity/service_auth/internal/handler"
	"github.com/AugustSerenity/service_auth/internal/service"
	"github.com/AugustSerenity/service_auth/internal/storage"
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
	s.ListenAndServe()
}
