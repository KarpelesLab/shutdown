package shutdown

import "fmt"

// errCh is a channel used to communicate fatal errors that should trigger shutdown
var errCh = make(chan error)

// Fatalf reports a fatal error that will trigger shutdown.
// It accepts format string with arguments similar to fmt.Errorf.
// It is possible to wrap errors using %w formatting directive.
func Fatalf(format string, args ...any) {
	errCh <- fmt.Errorf(format, args...)
}

// Go runs the provided function(s) in background goroutines.
// If any of these functions return a non-nil error, it will be handled as a fatal error
// and cause shutdown to trigger as if Fatalf was called.
// This is useful for starting background tasks that should cause the application
// to shut down gracefully if they encounter an error.
func Go(funcs ...func() error) {
	for _, f := range funcs {
		// Use a function parameter to correctly capture the loop variable
		go func(fn func() error) {
			err := fn()
			if err != nil {
				errCh <- err
			}
		}(f)
	}
}
