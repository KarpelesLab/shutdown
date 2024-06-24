package shutdown

import (
	"log"
	"slices"
)

var deferredFuncs []func()

// Defer will register the given function to be run upon program termination
func Defer(f func()) {
	deferredFuncs = append(deferredFuncs, f)
}

// runDefer executes the functions that were registered with Defer in reverse
// order.
func runDefer() {
	slices.Reverse(deferredFuncs)
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
