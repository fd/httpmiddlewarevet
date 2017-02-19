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
				handler := httpcache.NewHandler(httpcache.NewMemoryCache(), h)
				handler.Shared = true
				return handler
			},
		},
	)
}
