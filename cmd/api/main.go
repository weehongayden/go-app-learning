package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/weehongayden/bank-api/internal/config"
	"github.com/weehongayden/bank-api/internal/database"
	"github.com/weehongayden/bank-api/internal/logger"
)

func main() {
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log := logger.New(io.MultiWriter(os.Stdout, file), "", log.Ldate|log.Ltime)

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.LogError(log, fmt.Sprintf("Error loading config: %v", err))
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Pass, cfg.Database.Name)

	db, err := database.New(dsn)
	if err != nil {
		logger.LogError(log, fmt.Sprintf("Error connecting to database: %v", err))
	}

	svr := NewServer(log, cfg, db)

	err = svr.Start()
	if err != nil {
		logger.LogError(log, fmt.Sprintf("Error starting server: %v", err))
	}
}
