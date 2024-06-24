package shutdown

import "fmt"

var errCh = make(chan error)

// Fatalf reports an error, it is possible to wrap errors with %w
func Fatalf(format string, args ...any) {
	errCh <- fmt.Errorf(format, args...)
}
