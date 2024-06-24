package shutdown

import "log"

// Wait will wait until shutdown, or if an error is sent via Fatalf
func Wait() {
	select {
	case <-shutdownChannel:
	case err := <-errCh:
		log.Printf("[main] fatal error: %s", err)
	}
}
