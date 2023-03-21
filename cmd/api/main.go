package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/weehongayden/bank-api/internal/config"
	"github.com/weehongayden/bank-api/internal/database"
	"github.com/weehongayden/bank-api/internal/logger"
)

type App struct {
	logger *log.Logger
	config config.Config
	db     *sql.DB
}

func NewApp(log *log.Logger, config config.Config, db *sql.DB) *App {
	return &App{
		logger: log,
		config: config,
		db:     db,
	}
}

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

	app := NewApp(log, cfg, db)
	fmt.Println(app)
}
