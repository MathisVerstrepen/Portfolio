package portfolio

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func Init() {
	fmt.Println("Startup sequence starting...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("WARNING : Failed to load .env file")
	}
	fmt.Println("Startup sequence done.")
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func HTTPErrorHandler(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	if httpErr, ok := err.(*echo.HTTPError); ok {
		code = httpErr.Code
	}

	message := http.StatusText(code)
	if message == "" {
		code = http.StatusInternalServerError
		message = http.StatusText(code)
	}
	if code == http.StatusNotFound {
		message = "Page not found"
	}

	if code >= http.StatusInternalServerError {
		ctx.Logger().Error(err)
	}

	var responseErr error
	if ctx.Request().Method == http.MethodHead {
		responseErr = ctx.NoContent(code)
	} else {
		responseErr = ctx.String(code, message)
	}
	if responseErr != nil {
		ctx.Logger().Error(responseErr)
	}
}

func validLang(lang string) bool {
	return lang == "en" || lang == "fr"
}
