package shutdown

import (
	"fmt"
	"log/slog"
	"slices"
)

// deferredFuncs stores functions to be executed during shutdown in LIFO order
var deferredFuncs []func()

// Defer registers the given function to be run upon program termination.
// Functions registered with Defer will be executed in LIFO (Last In, First Out) order
// when shutdown is triggered.
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

// callDefer executes a deferred function and recovers from any panics.
// This ensures that all deferred functions will be called even if some panic.
func callDefer(f func()) {
	defer func() {
		if e := recover(); e != nil {
			slog.Error(fmt.Sprintf("panic in deferred shutdown function: %s (recovered)", e), "event", "shutdown:defer:panic", "category", "go.panic")
		}
	}()

	f()
}
