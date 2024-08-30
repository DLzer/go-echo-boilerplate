package main

import (
	"embed"
	"log"

	"github.com/DLzer/go-echo-boilerplate/internal/config"
	postgres "github.com/DLzer/go-echo-boilerplate/internal/database"
	"github.com/DLzer/go-echo-boilerplate/internal/server"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

// @title Echo Boilerplate
// @version 1.0
// @description Boilerplate code for an Echo API
// @contact.name DLzer
// @contact.url https://github.com/Dlzer
// @BasePath /v1
func main() {
	log.Println("=========Starting API Server==========")

	// Get config path
	cfgPath := utils.GetConfigPath("local")
	// Get config file
	cfgFile, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("LoadConfigError: %v", err)
	}

	// Parse config file
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfigError: %v", err)
	}

	// Starting Logger
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	// Postgres Connection
	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgres init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", "hello")
	}

	// Setup Embed and Postgres Dialect
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// Run Migrations
	db := stdlib.OpenDBFromPool(psqlDB)
	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	if err := db.Close(); err != nil {
		panic(err)
	}

	// Start the server
	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
