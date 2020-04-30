# gohealthcheck

Module that adds a health-checker to your Go app.

### Usage

Install the dependency

```
go get github.com/devkokov/gohealthcheck
```

Register the health-checker at the top of your main() function

```go
package main

import (
    "github.com/devkokov/gohealthcheck"
)

func main() {
    gohealthcheck.Register("http://localhost:8000/health-check")

    // ...
}
```

Run the health checker by adding a `-hc` flag to your executable

```
./your-go-build -hc
``` 
