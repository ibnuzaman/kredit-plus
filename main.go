package main

import (
	"context"
	"kredit-plus/config"
	"kredit-plus/database"
	"kredit-plus/internal"
	"kredit-plus/logger"
)

func init() {
	config.Init()
	ctx := context.Background()
	conf := config.Get()
	logger.Init(conf)
	database.Init(ctx, conf)
}

// @title						BE Kredit Plus
// @version					1.0
// @description				This is the API documentation for Kredit Plus backend services.
// @BasePath					/
// @securityDefinitions.apikey	AccessToken
// @in							header
// @name						Authorization
func main() {
	internal.Run()
}
