package system

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// NotFound
func NotFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "")
}
