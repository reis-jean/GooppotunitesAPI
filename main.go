package main

import (
	"github.com/reis-jean/GooppotunitesAPI.git/config"
	"github.com/reis-jean/GooppotunitesAPI.git/router"
)
var (
	logger *config.Logger
)

func main()  {
	
	logger =  config.GetLogger("main")

	err := config.Init()

    if err != nil {
        logger.ErrorF("Error initializing config: %v", err)
		return
    }

	router.Initialize()
}