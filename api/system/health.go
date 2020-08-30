package system

import (
	"github.com/labstack/echo/v4"
	"local/util"
	"net/http"
)

// Health health check
func Health(c echo.Context) error {
	var code int
	if !util.Health() {
		code = http.StatusInternalServerError
	} else {
		code = http.StatusOK
	}
	return c.String(code, "")
}
