package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// SetupSignals sets up signal handlers for graceful shutdown.
// It listens for interruptions (Ctrl-C/SIGINT or SIGTERM), and triggers a shutdown when
// a signal is received. If a second signal is sent (indicating shutdown is taking too long),
// this will force termination by calling os.Exit(1).
// This function should be called early in your application's main function.
func SetupSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		Shutdown()
		<-c
		os.Exit(1)
	}()
}
