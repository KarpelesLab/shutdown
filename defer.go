package shutdown

import "log"

var deferredFuncs []func()

// Defer will register the given function to be run upon program termination
func Defer(f func()) {
	deferredFuncs = append(deferredFuncs, f)
}

func runDefer() {
	for _, f := range deferredFuncs {
		callDefer(f)
	}
}

func callDefer(f func()) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic in deferred shutdown function: %s (recovered)", e)
		}
	}()

	f()
}
