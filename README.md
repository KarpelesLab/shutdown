[![GoDoc](https://godoc.org/github.com/KarpelesLab/shutdown?status.svg)](https://godoc.org/github.com/KarpelesLab/shutdown)

# shutdown

A lightweight Go library that provides utilities for graceful application shutdown handling, particularly useful for long-running services and daemons.

## Features

- **Signal Handling**: Automatically handles SIGINT (Ctrl-C) and SIGTERM signals
- **Graceful Shutdown**: Executes registered cleanup functions in LIFO order
- **Error Management**: Propagate fatal errors that should trigger application shutdown
- **Concurrency Support**: Run background goroutines that can trigger shutdown on error

## Installation

```
go get github.com/KarpelesLab/shutdown
```

## Usage 

```go
import (
    "net"
    "log/slog"

    "github.com/KarpelesLab/shutdown"
)

func main() {
    // Setup signal handlers early in your application
    shutdown.SetupSignals()

    // Run your application components
    go launchHttp()
    
    // Example of running multiple background tasks with error handling
    shutdown.Go(
        task1,
        task2,
        // More tasks...
    )

    // Wait for shutdown (this should be the last call in main)
    shutdown.Wait()
}

func launchHttp() {
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        // Report fatal errors that should trigger shutdown
        shutdown.Fatalf("failed to listen for the http server: %w", err)
        return
    }

    // Register cleanup functions to be executed during shutdown
    shutdown.Defer(func() {
        slog.Info("Closing HTTP listener")
        l.Close()
    })
    
    // Your HTTP server implementation...
}

// Example task function that returns an error
func task1() error {
    // If this returns an error, it will trigger shutdown
    return nil
}
```

## Core Functions

- `SetupSignals()`: Configures signal handlers for graceful shutdown (call early in main)
- `Defer(func())`: Registers cleanup functions to be executed during shutdown (LIFO order)
- `Fatalf(format, args...)`: Reports a fatal error that triggers shutdown
- `Go(funcs ...func() error)`: Runs functions in background goroutines, errors trigger shutdown
- `Shutdown()`: Manually triggers the shutdown process
- `Wait()`: Blocks until shutdown is triggered, then runs cleanup functions (call last in main)

## Shutdown Triggers

Shutdown can be triggered in several ways:

1. OS Signals (SIGINT/Ctrl-C or SIGTERM)
2. Fatal errors reported via `Fatalf()`
3. Errors returned from functions started with `Go()`
4. Manual trigger via `Shutdown()`

When shutdown is triggered, all functions registered with `Defer()` are executed in LIFO order (last registered, first executed).

## License

See [LICENSE](LICENSE) file.