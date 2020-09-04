package main

import (
	"fmt"
	apiSystem "local/api/system"
	apiUser "local/api/user"
	mw "local/middleware"
	"local/util"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(mw.RestLogger)

	e.GET("/health", apiSystem.Health)
	e.GET("/user", apiUser.Test)

	//e.Use(mw.LoggerPost)
	e.Any("/*", apiSystem.NotFound)

	e.Logger.Fatal(e.StartServer(&http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}))
}
