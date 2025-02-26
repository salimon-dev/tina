package main

import (
	"salimon/tina-core/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
	}))
	// HTTP route
	e.GET("/", rest.HeartBeatHandler)
	// WebSocket route
	e.GET("/interact", rest.InteractHandler)
	// Start the server
	port := "80"
	e.Logger.Fatal(e.Start(":" + port))
}
