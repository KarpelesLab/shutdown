package shutdown

import "log"

var shutdownChannel = make(chan struct{})

func shutdown() {
	log.Println("[main] shutting down...")
	close(shutdownChannel)
}
