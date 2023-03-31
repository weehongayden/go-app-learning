package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weehongayden/bank-api/internal/config"
)

type Server struct {
	logger *log.Logger
	config config.Config
	db     *sql.DB
}

const (
	defaultIdleTimeout  = time.Minute
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

func NewServer(logger *log.Logger, config config.Config, db *sql.DB) *Server {
	return &Server{
		logger: logger,
		config: config,
		db:     db,
	}
}

func (app Server) Start() error {
	r := NewRoutes(*app.logger, *app.db)

	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.config.Server.Host, app.config.Server.Port),
		Handler:      r.Routes(),
		ErrorLog:     app.logger,
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	err := s.ListenAndServe()
	if err != nil {
		app.logger.Fatalf("Error starting server: %v", err)
		return err
	}

	app.logger.Printf("Server started on %s:%d", app.config.Server.Host, app.config.Server.Port)

	return nil
}
