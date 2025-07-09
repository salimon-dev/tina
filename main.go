package main

import (
	"fmt"
	"os"
	"salimon/tina/db"
	"salimon/tina/openai"
	"salimon/tina/webhook"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("no environment file, using session defaults")
	}
	db.SetupDatabase()
	openai.LoadActions()
	e := echo.New()
	e.HideBanner = true
	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
	}))
	// heartbeat route to check if server is alive
	e.GET("/", HeartBeatHandler)
	// interact main route for user interaction
	e.POST("/webhook", webhook.Handle)
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
