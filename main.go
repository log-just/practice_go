package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apiSystem "local/api/system"
	apiUser "local/api/user"
	mw "local/middleware"
	"local/util"
	"os"
)

func main() {

	// init internal libs
	if !util.Init() {
		fmt.Println("fail to init")
		//debug.PrintStack()
		os.Exit(1)
	}

	// init web
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.Logger)

	e.GET("/health", apiSystem.Health)
	e.GET("/user", apiUser.Test)

	//e.Use(mw.LoggerPost)
	e.Any("/*",apiSystem.NotFound)

	e.Logger.Fatal(e.Start(":1323"))
}
