package shutdown

import (
	"fmt"
	"log/slog"
)

// Wait blocks until shutdown is triggered, either by calling Shutdown(),
// receiving a termination signal (if SetupSignals was called), or if an error is sent via Fatalf.
// After the shutdown is triggered, this function runs all cleanup functions registered with Defer().
// This function should be called as the last function in main().
func Wait() {
	select {
	case <-shutdownChannel:
	case err := <-errCh:
		if err != nil {
			slog.Info(fmt.Sprintf("[shutdown] Shutting down on fatal error: %s", err), "event", "shutdown:fatal")
		}
	}

	RunDefer()
}
