package shutdown

import "log"

var shutdownChannel = make(chan struct{})

// Shutdown will trigger the normal shutdown of the program
func Shutdown() {
	log.Println("[shutdown] shutting down...")
	close(shutdownChannel)
}
