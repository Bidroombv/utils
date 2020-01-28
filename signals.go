package utils

import (
	"os"
	"os/signal"
	"syscall"
)

// InterceptTerminationSignals intercepts termination signals, and notifies the
// caller by writing to the shutdown channel
func InterceptTerminationSignals(shutdown chan<- bool) {
	signals := make(chan os.Signal, 1)

	// catch main termination signals from OS
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGABRT)

	go func() {
		<-signals
		shutdown <- true
	}()
}
