# newrelic-handler

New Relic agent handler for Go's net/http package.

## Usage

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/mikoim/newrelic-handler"
)

func main() {
	// Routing
	router := httprouter.New()
	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	// New Relic
	n, _ := nrh.New(nrh.Options{
		ApplicationName: os.Getenv("NEWRELIC_APPNAME"),
		LicenseKey:      os.Getenv("NEWRELIC_LICENSE"),
	})
	handler := n.Handler(router)

	http.ListenAndServe(":8080", handler)
}
```