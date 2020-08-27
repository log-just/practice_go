package main

import (
	"github.com/labstack/echo/v4"

	"local/route"
)

func main() {
	e := echo.New()
	route.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
