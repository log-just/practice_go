package system

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NotFound : rest api가 없을때 restLogger가 인지할수있도록 에러 발생
func NotFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "")
}
