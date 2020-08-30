package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// validator 인스턴스
var val = validator.New()

// Validate : struct 검증
func Validate(class interface{}) error {
	return val.Struct(class)
}

// ValidateReq : REST API 요청 body 검증
func ValidateReq(c echo.Context, class interface{}) error {
	if err := c.Bind(class); err != nil {
		return err
	}
	if err := val.Struct(class); err != nil {
		return err
	}
	return nil
}
