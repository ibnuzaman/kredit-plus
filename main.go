package main

import (
	"context"
	"kredit-plus/config"
	"kredit-plus/database"
	"kredit-plus/logger"
)

func init() {
	config.Init()
	ctx := context.Background()
	conf := config.Get()
	logger.Init(conf)
	database.Init(ctx, conf)
}

func main() {
	log := logger.Get("main")
	log.Info().Msgf("Starting Kredit Plus Service with config: %+v", config.Get())

	db := database.Get()
	log.Info().Msgf("Database connection established: %v", db != nil)
}
