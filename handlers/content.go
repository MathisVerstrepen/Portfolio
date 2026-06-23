package portfolio

import (
	"net/http"

	"github.com/labstack/echo/v4"

	comp "portfolio/components"
)

func ServeTechAndAcademicsPage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.TechAndAcademics(lang), "Tech & Academics", lang))
}

func ServeExperiencePage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.Experience(lang), "Experience", lang))
}

func ServeProjectsPage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.Projects(lang), "Projects", lang))
}

func ServeContactPage(c echo.Context) error {
	return Render(c, http.StatusOK, comp.Root(comp.Contact(), "Contact", "en"))
}
