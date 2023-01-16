package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// healthcheck responds to a healthcheck request.
func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// RegisterHandlers registers the handlers that perform healthchecks.
func Handle(e *echo.Echo) {
	e.GET("", healthcheck)
}
