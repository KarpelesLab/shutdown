package shutdown

import "log"

// Wait will wait until shutdown, or if an error is sent via Fatalf
func Wait() {
	select {
	case <-shutdownChannel:
	case err := <-errCh:
		if err != nil {
			log.Printf("[main] fatal error: %s", err)
		}
	}

	runDefer()
}
