# HTTP Middleware Vet

---

#### Validate Middleware Handlers with Vet

Many http.Handler wrappers (e.g. middleware) do not implement all needed interfaces. **Vet** provides a test to validate these handlers.

example :

```go
package main

import (
	"net/http"
	"testing"

	httpmiddlewarevet "github.com/fd/httpmiddlewarevet"
)

func Test(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return MyMiddleware(h) // replace MyMiddleware with your http.Handler
	})
}
```

----

#### Must Implement :

- HTTP/1.1
  - http.Flusher
  - http.Hijacker
  - http.CloseNotifier
  - io.ReaderFrom
- TLS/HTTP/1.1
  - TLS
  - http.Flusher
  - http.Hijacker
  - http.CloseNotifier
  - io.ReaderFrom
- TLS/HTTP/2.0
  - TLS
  - http.Flusher
  - http.CloseNotifier
  - http.Pusher
