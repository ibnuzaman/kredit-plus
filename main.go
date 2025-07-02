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

func main() {
	internal.Run()
}
