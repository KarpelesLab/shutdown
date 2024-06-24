# shutdown

Utility library to handle daemons

## Usage 

```go
import "github.com/KarpelesLab/shutdown"

func init() {
    shutdown.SetupSignals()

    // do the things you want to do
    go launchHttp()

    shutdown.Wait()
}
```
