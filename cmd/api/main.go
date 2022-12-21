package main

import (
	"context"
	"log"
	"os"

	"github.com/DLzer/go-echo-boilerplate/config"
	"github.com/DLzer/go-echo-boilerplate/internal/server"
	mongodb "github.com/DLzer/go-echo-boilerplate/pkg/db/mongodb"
	"github.com/DLzer/go-echo-boilerplate/pkg/db/postgres"
	"github.com/DLzer/go-echo-boilerplate/pkg/db/s3"
	"github.com/DLzer/go-echo-boilerplate/pkg/logger"
	"github.com/DLzer/go-echo-boilerplate/pkg/utils"
)

// @title Echo Boilerplate
// @version 1.0
// @description Boilerplate code for an Echo API
// @contact.name DLzer
// @contact.url https://github.com/Dlzer
// @BasePath /v1
func main() {
	log.Println("=========Starting API Server==========")

	// Loading config
	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	// Parse ( Unmarshal ) config
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
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
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}

	// Mongo Connection
	mongo, err := mongodb.NewMongoDB(cfg)
	if err != nil {
		appLogger.Fatalf("Mongo Init: %s", err)
	}
	appLogger.Info("Mongo Connected")
	// Defer closing the mongo connection for re-use
	defer func() {
		if err = mongo.MongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// S3 Connection
	s3Client, err := s3.NewS3Client(cfg.S3.SpacesKey, cfg.S3.SpacesSecret, cfg.S3.SpacesEndpoint, cfg.S3.SpacesRegion)
	if err != nil {
		appLogger.Fatalf("S3 init: %s", err)
	} else {
		appLogger.Info("S3 Connected")
	}

	// Start the server
	s := server.NewServer(cfg, psqlDB, mongo, s3Client, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
