package shutdown

import (
	"fmt"
	"log/slog"
)

// Wait will wait until shutdown, or if an error is sent via Fatalf. This function
// should be called as the last function in main()
func Wait() {
	select {
	case <-shutdownChannel:
	case err := <-errCh:
		if err != nil {
			slog.Info(fmt.Sprintf("[shutdown] Shutting down on fatal error: %s", err), "event", "shutdown:fatal")
		}
	}

	runDefer()
}
