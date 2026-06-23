package portfolio

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func publicSiteURL() string {
	url := strings.TrimRight(os.Getenv("SITE_URL"), "/")
	if url == "" {
		return "https://portfolio.diikstra.fr"
	}
	return url
}

func ServeRobots(c echo.Context) error {
	body := "User-agent: *\nAllow: /\nSitemap: " + publicSiteURL() + "/sitemap.xml\n"
	return c.Blob(http.StatusOK, "text/plain; charset=utf-8", []byte(body))
}

func ServeSitemap(c echo.Context) error {
	paths := []string{
		"/en",
		"/fr",
		"/en/tech&academics",
		"/fr/tech&academics",
		"/en/experience",
		"/fr/experience",
		"/en/projects",
		"/fr/projects",
		"/contact",
	}

	var body strings.Builder
	body.WriteString(xml.Header)
	body.WriteString("<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n")
	for _, path := range paths {
		body.WriteString("  <url><loc>")
		body.Write(escapedXML(publicSiteURL() + path))
		body.WriteString("</loc></url>\n")
	}
	body.WriteString("</urlset>\n")

	return c.Blob(http.StatusOK, "application/xml; charset=utf-8", []byte(body.String()))
}

func escapedXML(value string) []byte {
	var result bytes.Buffer
	_ = xml.EscapeText(&result, []byte(value))
	return result.Bytes()
}
