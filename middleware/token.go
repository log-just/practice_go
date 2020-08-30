package middleware

import (
	"github.com/labstack/echo/v4"
	"local/util"
	"net/http"
)

// CheckToken : 유저의 jwt 토큰을 확인해보자~
func CheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// get jwt Token
		jwtToken := c.Request().Header.Get("tkn")
		if jwtToken == "" {
			return c.String(http.StatusUnauthorized, "")
		}

		// verify & get Data
		data, err := util.JwtVerify(jwtToken, c.RealIP())
		if err != nil {
			return c.String(http.StatusUnauthorized, "")
		}
		// set datas to Context
		c.Set("id", data.ID)

		return next(c)

	}
}
