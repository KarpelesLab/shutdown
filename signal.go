package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// SetupSignals will listen for interruptions (Ctrl-C or kill), and trigger a shutdown if
// a signal is received. If a second signal is sent (shutdown is taking too long), this will
// call os.Exit(1)
func SetupSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		shutdown()
		<-c
		os.Exit(1)
	}()
}
