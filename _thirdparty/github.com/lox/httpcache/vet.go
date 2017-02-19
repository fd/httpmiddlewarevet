package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/lox/httpcache"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "NewHandler",
			Func: func(h http.Handler) http.Handler {
				return httpcache.NewHandler(httpcache.NewMemoryCache(), h)
			},
		},
	)
}
