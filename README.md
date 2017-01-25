[![Build Status](https://travis-ci.org/fd/httpmiddlewarevet.svg?branch=master)](https://travis-ci.org/fd/httpmiddlewarevet)
[![GoDoc](https://godoc.org/github.com/fd/httpmiddlewarevet?status.svg)](https://godoc.org/github.com/fd/httpmiddlewarevet)

#### Build Status

We run daily tests on commonly used third party middleware. (see `/_thirdparty`)
If builds fail checkout Travis CI to find out which packages are affected.
We test against go1.5, go1.6, go1.7, go master.

# HTTP Middleware Vet

#### Validate Middleware Handlers with Vet

Many http.Handler wrappers (e.g. middleware) do not implement all needed
interfaces or respect basic http/x semantics.
**Vet** provides a test to validate these handlers.

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

### What is tested for

* `http.Flusher`
  * `go1.x`: interface compliance
* `http.Hijacker`
  * `go1.x`: interface compliance
* `http.CloseNotifier`
  * `go1.x`: interface compliance
  * `go1.x`: semantics
* `io.ReaderFrom`
  * `go1.x`: interface compliance
* `http.Pusher`
  * `go1.8+`: interface compliance
* `http.Request.Context`
  * `go1.x`: absent
  * `go1.8+`: Done() semantics
