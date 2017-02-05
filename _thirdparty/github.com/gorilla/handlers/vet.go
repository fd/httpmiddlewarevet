package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/gorilla/handlers"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "CompressHandler",
			Func: func(h http.Handler) http.Handler {
				return handlers.CompressHandler(h)
			},
		},
		testing.Middleware{
			Name: "CompressHandlerLevel",
			Func: func(h http.Handler) http.Handler {
				return handlers.CompressHandlerLevel(h, 4)
			},
		},
	)
}
