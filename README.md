[![Build Status](https://travis-ci.org/fd/httpmiddlewarevet.svg?branch=master)](https://travis-ci.org/fd/httpmiddlewarevet)
[![GoDoc](https://godoc.org/github.com/fd/httpmiddlewarevet?status.svg)](https://godoc.org/github.com/fd/httpmiddlewarevet)

#### Build Status

We run daily tests on commonly used third party middleware. (see `/thirdparty`) If builds fail checkout Travis CI to find out which packages are affected.
We test against go1.6, go1.7, go latest and go master.

# HTTP Middleware Vet

#### Validate Middleware Handlers with Vet

Many http.Handler wrappers (e.g. middleware) do not implement all needed interfaces. **Vet** provides a test to validate these handlers.

example :

```go
package main

import (
	"net/http"
	"testing"

	"github.com/fd/httpmiddlewarevet"
)

func Test(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return MyMiddleware(h) // replace MyMiddleware with your http.Handler
	})
}
```

----

#### Middleware Must Implement :

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
