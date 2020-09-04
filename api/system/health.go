package system

import (
	"local/util"
	"net/http"

	"github.com/labstack/echo/v4"
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
