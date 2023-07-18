package routes

import (
	"now-iusearchbtw/config"
	"now-iusearchbtw/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, c *config.Config) {
	e.Static("/", "public")
	e.GET("/ping", controllers.Ping())
	e.GET("/new", controllers.NewContainer(c))
	e.DELETE("/kill", controllers.KillContainer(c))
}
