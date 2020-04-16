package utils

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

// Log is the logger
var Log *Logger

func (l Logger) Verbose() bool {
	return true // TODO
}

// LoggingFunc is a type of logging function which returns a true boolean in case of error
type LoggingFunc func(error, string, ...interface{}) bool

// LogSetup initialize the logger with proper server name and flags
func LogSetup(serverName string) {
	srv := "[" + serverName + "] "
	Log = &Logger{
		Logger: log.New(os.Stderr, srv, log.LstdFlags|log.Lmicroseconds|log.LUTC|log.Llongfile),
	}
}

// FailOnError logs and fails in case of err is != nil
func FailOnError(err error, format string, a ...interface{}) {
	if err != nil {
		// Ignore logging error. What else can we do, log it?
		_ = Log.Output(2, fmt.Sprintf("%s: %s", fmt.Sprintf(format, a...), err))
		os.Exit(1)
	}
}

// LogOnError logs and returns true in case of err is != nil
func LogOnError(err error, format string, a ...interface{}) bool {
	if err != nil {
		// Ignore logging error. What else can we do, log it?
		_ = Log.Output(2, fmt.Sprintf("%s: %s", fmt.Sprintf(format, a...), err))
		return true
	}
	return false
}

// CorrelatedLogError logs errors with a correlation ID
func CorrelatedLogError(cid string, logFunc LoggingFunc) LoggingFunc {
	return func(err error, format string, a ...interface{}) bool {
		origMsg := fmt.Sprintf(format, a...)
		augmentedMsg := fmt.Sprintf("%s on Message with CorrelationId: %s", origMsg, cid)

		return logFunc(err, augmentedMsg)
	}
}
