package system

import (
	"github.com/labstack/echo/v4"
	"local/util"
	"net/http"
)
// Health health check
func Health(c echo.Context) error {
	println("api")
	if isHealth:= util.Health();isHealth !=true {
		return c.String(http.StatusInternalServerError, "")
	}
	return c.String(http.StatusOK, "")
}
