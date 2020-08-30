package util

import (
	"cloud.google.com/go/logging"
	"context"
	"fmt"
)

var googleLogger *logging.Logger

func logInit() bool {
	ctx := context.Background()
	projectID := "YOUR_PROJECT_ID"
	//var err error
	googleLogging, err := logging.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("fail to init logger : %v",err)
		return true
		//return false
	}
	logName := "my-log"
	googleLogger = googleLogging.Logger(logName)
	return true
}

func LogInfo(log interface{})  {
	googleLogger.Log(logging.Entry{Severity: logging.Info, Payload: log})
}

func LogErr(log ErrResponse)  {
	googleLogger.Log(logging.Entry{Severity: logging.Error, Payload: log})
}

