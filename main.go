package main

import (
	"github.com/joaodslourenco/gopportunities/config"
	"github.com/joaodslourenco/gopportunities/router"
)

var logger *config.Logger

func main() {

	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
	}

	router.Initialize()
}
