package user

import (
	"local/route"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get : 라우팅 설정!
func Get(c *echo.Context, p *route.ReqUser) {
	return c.JSON(http.StatusOK, p)
}
