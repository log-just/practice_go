package route

import (
	"github.com/labstack/echo/v4"
)

// Init : 라우팅 설정!
func Init(e *echo.Echo) {
	initUser(e)
}
