package user

import (
	"local/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bodyReq struct {
	Name int `query:"name" validate:"required"`
	Tame int `json:"tame" form:"tame" query:"tame" validate:"required"`
}

type bodyRes struct {
	Name string `json:"name" form:"name" query:"name"`
}

// Test : /user
func Test(c echo.Context) error {
	p := new(bodyReq)
	if err := util.ValidateReq(c, p); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrParamBodyWithVerbose(err))
	}
	return c.JSON(http.StatusOK, p)
}
