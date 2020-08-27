package route

import (
	"net/http"

	"local/api/user"
	"local/util"

	"github.com/labstack/echo/v4"
)

func getErrParamBody(err error) *util.ErrResponse {
	return &util.ErrResponse{
		Code:   util.ErrCodeParam,
		Detail: err.Error(),
	}
}

// ReqUser : 라우팅 설정!
type ReqUser struct {
	Name string `json:"name"`
}

// ResUser : 라우팅 설정!
type ResUser struct {
	Name string `json:"name" form:"name" query:"name"`
}

func initUser(e *echo.Echo) {
	e.GET("/user", func(c echo.Context) error {
		p := new(ReqUser)
		if err := c.Bind(p); err != nil {
			return c.JSON(http.StatusBadRequest, getErrParamBody(err))
		}
		return user.Get(p)
	})
}
