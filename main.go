package main

import (
	"fmt"
	"os"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handlers "portfolio/handlers"
)

func main() {
	e := echo.New()
	if os.Getenv("ENV") == "dev" {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/assets", path.Join(os.Getenv("ROOT_DIR"), "assets"))

	handlers.Init()

	// ---- Home Routes ---- //
	e.GET("/", handlers.ServeHomePage)
	e.GET("/:lang", handlers.ServeHomePage)

	// ---- Global Routes ---- //
	e.GET("/ping", handlers.ServePing)
	if os.Getenv("ENV") != "prod" {
		e.GET("/ws", handlers.InitHotReloadWebSocket)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
