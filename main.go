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
	handlers.Init()

	if os.Getenv("ENV") == "dev" {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/assets", path.Join(os.Getenv("ROOT_DIR"), "assets"))

	// ---- Home Routes ---- //
	e.GET("/", handlers.ServeHomePage)
	e.GET("/:lang", handlers.ServeHomePage)

	// ---- Content Routes ---- //
	e.GET("/:lang/tech&academics", handlers.ServeTechAndAcademicsPage)
	e.GET("/:lang/experience", handlers.ServeExperiencePage)
	e.GET("/:lang/projects", handlers.ServeProjectsPage)
	e.GET("/contact", handlers.ServeContactPage)

	// ---- Global Routes ---- //
	e.GET("/ping", handlers.ServePing)
	if os.Getenv("ENV") != "prod" {
		e.GET("/ws", handlers.InitHotReloadWebSocket)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
