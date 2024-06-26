[![GoDoc](https://godoc.org/github.com/KarpelesLab/shutdown?status.svg)](https://godoc.org/github.com/KarpelesLab/shutdown)

# shutdown

Utility library to handle daemons

## Usage 

```go
import (
    "net"

    "github.com/KarpelesLab/shutdown"
)

func main() {
    shutdown.SetupSignals()

    // do the things you want to do
    go launchHttp()

    shutdown.Wait()
}

func launchHttp() {
    l, err := net.Listen("tcp", ":80")
    if err != nil {
        shutdown.Fatalf("failed to listen for the http server: %w", err)
        return
    }

    // cleanup of opened listen socket
    shutdown.Defer(func() {
        l.Close()
    })
    // ...
}
```
