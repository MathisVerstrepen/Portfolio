package portfolio

import (
	"net/http"

	"github.com/labstack/echo/v4"

	comp "portfolio/components"
)

func ServeHomePage(c echo.Context) error {
	lang := c.Param("lang")
	if lang == "" {
		lang = "en"
	}

	return Render(c, http.StatusOK, comp.Root(comp.Home(lang), "Home"))
}
