package main

import (
	"kredit-plus/config"
	"kredit-plus/logger"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.Init(conf)
}

func main() {
	log := logger.Get("main")
	log.Info().Msgf("Starting Kredit Plus Service with config: %+v", config.Get())
}
