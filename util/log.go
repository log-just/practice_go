package util

import (
	"context"
	"fmt"
	"runtime/debug"

	"cloud.google.com/go/logging"
)

var googleLogger *logging.Logger

// LogTypeMetric 메트릭용 로그타입
type LogTypeMetric struct {
	Typ string
	Nam string
}

func initLog() error {
	ctx := context.Background()
	projectID := "YOUR_PROJECT_ID"
	//var err error
	googleLogging, err := logging.NewClient(ctx, projectID)
	if err != nil {
		fmt.Printf("fail to init logger : %v", err)
		debug.PrintStack()
		return err
	}
	logName := "my-log"
	googleLogger = googleLogging.Logger(logName)
	return nil
}

// LogInfo : info 레벨 로그
func LogInfo(log interface{}) {
	googleLogger.Log(logging.Entry{Severity: logging.Info, Payload: log})
}

// LogErr : error 레벨 로그
func LogErr(log ErrResponse) {
	googleLogger.Log(logging.Entry{Severity: logging.Error, Payload: log})
}
