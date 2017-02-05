[![Build Status](https://travis-ci.org/fd/httpmiddlewarevet.svg?branch=master)](https://travis-ci.org/fd/httpmiddlewarevet)
[![GoDoc](https://godoc.org/github.com/fd/httpmiddlewarevet?status.svg)](https://godoc.org/github.com/fd/httpmiddlewarevet)

Finding high quality `net/http` middleware is hard. Middleware get outdated,
have only partial support or fail to comply with basic `net/http` semantics
in other ways. The `httpmiddlewarevet` project aims to provide developers
with a clearer picture of which handlers are good `net/http` citizens and which
are not.

#### Build Status

We run daily tests on commonly used third party middleware. (see `/_thirdparty`)
If builds fail checkout Travis CI to find out which packages are affected.
We test against go1.5, go1.6, go1.7, go master.

# HTTP Middleware Vet

#### Validate Middleware Handlers with Vet

Many `http.Handler` wrappers (e.g. middleware) do not implement all necessary
interfaces or respect the basic semantics of those interfaces. **Vet** provides
a tool to test these handlers.

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

This list is expected to grow over time. Contributions are welcome.


### Which version of Go are tested

The last stable release and the two previous minor releases as well as
tip are tested. So if `go1.7` is the latest stable release `go1.7`, `go1.6`,
`go1.5` and `go1.8-pre` are tested. This gives any middleware maintainers the
leeway to guarantee _stable && stable-1_ stability.
