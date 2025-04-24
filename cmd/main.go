package main

import (
	"net/http"

	"github.com/AugustSerenity/service_auth/internal/handler"
	"github.com/AugustSerenity/service_auth/internal/service"
	"github.com/AugustSerenity/service_auth/internal/storage"
)

const portNumber = ":8080"

func main() {

	db := storage.InitDB()
	defer storage.CloseDB(db)

	storage := storage.New(db)

	srv := service.New(storage)

	h := handler.New(srv)

	s := http.Server{
		Addr:    portNumber,
		Handler: h.Route(),
	}

	s.ListenAndServe()
}
