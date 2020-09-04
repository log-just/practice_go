package middleware

import (
	"fmt"
	"local/util"
	"time"

	"github.com/labstack/echo/v4"
)

type logRes struct {
	util.LogTypeMetric
	req string
	lat int64
	res int
}

// RestLogger : rest api 로깅
func RestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("t", time.Now())

		err := next(c)

		resCode := c.Response().Status
		req := c.Request()
		var url string
		if err != nil || resCode == 404 {
			url = "other"
		} else {
			url = req.URL.Path
		}
		startTime, _ := c.Get("t").(time.Time)
		//util.LogInfo(&logRes{
		//	typ: "Mrest",
		//	req: fmt.Sprintf("%s %s",req.Method,url),
		//	lat: time.Since(startTime).Seconds(),
		//	res: resCode,
		//})

		fmt.Printf("%+v\n", &logRes{
			util.LogTypeMetric{
				Typ: "metric",
				Nam: "rest",
			},
			fmt.Sprintf("%s %s", req.Method, url),
			time.Since(startTime).Milliseconds(),
			resCode,
		})
		return err
	}
}

//// Logger 로깅
//func LoggerPost(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		println("post")
//		return next(c)
//	}
//}
