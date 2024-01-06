package main

import (
	"Api_first_1_19/config"
	"Api_first_1_19/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
