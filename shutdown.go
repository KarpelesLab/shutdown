// Package shutdown provides utilities for graceful application shutdown handling.
package shutdown

import (
	"log/slog"
	"sync"
)

var (
	// shutdownChannel is closed when shutdown is triggered
	shutdownChannel = make(chan struct{})
	// shutdownOnce ensures shutdown logic runs exactly once
	shutdownOnce sync.Once
)

// Shutdown will trigger the normal shutdown of the program.
// It can be called multiple times safely as it uses sync.Once internally.
func Shutdown() {
	shutdownOnce.Do(func() {
		slog.Info("[shutdown] Shutdown requested", "event", "shutdown:request")
		close(shutdownChannel)
	})
}
