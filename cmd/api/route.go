package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type IRoute[T any] interface {
	Register(r chi.Router, args T)
	All()
	Get()
	Post()
	Put()
	Delete()
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type Route struct {
	logger log.Logger
	db     sql.DB
}

func NewRoutes(logger log.Logger, db sql.DB) *Route {
	return &Route{
		logger: logger,
		db:     db,
	}
}

func (app Route) Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	RegisterCardRoute(router, app)

	router.Mount("/api/v1", router)

	return router
}
