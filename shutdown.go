package shutdown

import (
	"log/slog"
	"sync"
)

var (
	shutdownChannel = make(chan struct{})
	shutdownOnce    sync.Once
)

// Shutdown will trigger the normal shutdown of the program
func Shutdown() {
	shutdownOnce.Do(func() {
		slog.Info("[shutdown] Shutdown requested", "event", "shutdown:request")
		close(shutdownChannel)
	})
}
