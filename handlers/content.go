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

	return Render(c, http.StatusOK, comp.Root(comp.TechAndAcademics(lang), "Tech & Academics"))
}

func ServeExperiencePage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.Experience(lang), "Experience"))
}

func ServeProjectsPage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.Projects(lang), "Projects"))
}
