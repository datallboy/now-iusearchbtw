package main

import (
	"fmt"
	"log"
	"now-iusearchbtw/config"
	"now-iusearchbtw/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config, err := config.New()
	if err != nil {
		log.Fatal("Error creating configuration", err)
	}

	routes.Routes(e, config)

	address := fmt.Sprintf("%s:%s", config.ListeningAddress, config.ListeningPort)

	e.Logger.Fatal(e.Start(address))
}
