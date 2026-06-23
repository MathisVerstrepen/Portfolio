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
	e.HTTPErrorHandler = handlers.HTTPErrorHandler
	handlers.Init()

	if os.Getenv("ENV") == "dev" {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "DENY",
		ReferrerPolicy:     "strict-origin-when-cross-origin",
		ContentSecurityPolicy: "default-src 'self'; " +
			"base-uri 'self'; " +
			"object-src 'none'; " +
			"frame-ancestors 'none'; " +
			"img-src 'self' data:; " +
			"font-src 'self' https://fonts.gstatic.com; " +
			"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; " +
			"script-src 'self' 'unsafe-inline'; " +
			"connect-src 'self' ws: wss:; " +
			"form-action 'self'",
	}))
	e.Use(permissionsPolicyHeader)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Static("/assets", path.Join(os.Getenv("ROOT_DIR"), "assets"))

	// ---- Home Routes ---- //
	e.GET("/", handlers.ServeHomePage)
	e.GET("/:lang", handlers.ServeHomePage)

	// ---- Content Routes ---- //
	e.GET("/:lang/tech-academics", handlers.ServeTechAndAcademicsPage)
	e.GET("/:lang/experience", handlers.ServeExperiencePage)
	e.GET("/:lang/projects", handlers.ServeProjectsPage)
	e.GET("/contact", handlers.ServeContactPage)

	// ---- Legacy Redirects ---- //
	e.GET("/:lang/tech&academics", handlers.RedirectLegacyTechAcademics)

	// ---- Global Routes ---- //
	e.GET("/robots.txt", handlers.ServeRobots)
	e.GET("/sitemap.xml", handlers.ServeSitemap)
	e.GET("/ping", handlers.ServePing)
	if os.Getenv("ENV") != "prod" {
		e.GET("/ws", handlers.InitHotReloadWebSocket)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func permissionsPolicyHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=(), payment=(), usb=(), browsing-topics=()")
		return next(c)
	}
}
