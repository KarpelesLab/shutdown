package shutdown

import "log"

// Wait will wait until shutdown, or if an error is sent via Fatalf. This function
// should be called as the last function in main()
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
