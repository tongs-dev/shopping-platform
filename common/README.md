# Common Package

The `common` package provides shared utilities and configurations for the application. 
It includes logging, database connectivity, observability, and other common components used across services.

## 📂 Directory Structure
```
common/ 
    │── config.go     # Configuration management (Consul, environment variables)
    │── jaeger.go     # Jaeger tracing setup
    │── mysql.go      # MySQL database connection setup
    │── prometheus.go # Prometheus metrics configuration
    │── swap.go       # Utility for struct conversion using JSON tags
    │── zap.go        # Logging setup using Zap
    │── go.mod        # Module definition for dependency management
    │── go.sum        # Dependency checksums
    │── README.md     # Documentation
```

## 🚀 **Setup Instructions**
1. Install Dependencies
```sh
go mod tidy # Ensure all dependencies are installed
```

2. Load Configuration
```go
import "your_project/common"

conf := common.GetConfig()
```

3. Initialize Logging
```go
import "your_project/common"

logger := common.InitLogger()
logger.Info("Application started")
```

4. Connect to MySQL
```go
import "your_project/common"

db := common.GetMySQLConnection()
```

5. Enable Jaeger Tracing
```go
import "your_project/common"

tracer, closer, err := common.NewTracer("my-service", "localhost:6831")
defer closer.Close()

```

6. Start Prometheus Metrics
```go
import "your_project/common"

common.PrometheusBoot(9090) // Exposes metrics on /metrics endpoint
```

7. Struct Conversion Utility
```go
import "your_project/common"

type Source struct {
    Name string `json:"name"`
}

type Target struct {
    Name string `json:"name"`
}

src := Source{Name: "Example"}
var tgt Target

if err := common.SwapTo(src, &tgt); err != nil {
    log.Fatal(err)
}
```
