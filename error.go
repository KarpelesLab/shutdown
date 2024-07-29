package shutdown

import "fmt"

var errCh = make(chan error)

// Fatalf reports an error, it is possible to wrap errors with %w
func Fatalf(format string, args ...any) {
	errCh <- fmt.Errorf(format, args...)
}

// Go runs function(s) f in background, and if any of those functions
// return an error, it'll be handled as a fatal error and cause shutdown
// to trigger as if Fatalf was called
func Go(funcs ...func() error) {
	for _, f := range funcs {
		go func() {
			err := f()
			if err != nil {
				errCh <- err
			}
		}()
	}
}
