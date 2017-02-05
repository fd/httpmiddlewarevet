package main

import (
	"net/http"
	"net/url"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/vulcand/oxy/forward"
)

func main() {

	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		req.URL, _ = url.Parse("http://localhost:63450")
		fwd.ServeHTTP(w, req)
	})

	testing.Run(
		testing.Middleware{
			Name: "forward.ServeHTTP",
			Func: func(h http.Handler) http.Handler {
				return redirect
			},
		},
	)
}
