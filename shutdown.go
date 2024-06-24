package shutdown

import "log/slog"

var shutdownChannel = make(chan struct{})

// Shutdown will trigger the normal shutdown of the program
func Shutdown() {
	slog.Info("[shutdown] Shutdown requested", "event", "shutdown:request")
	close(shutdownChannel)
}
