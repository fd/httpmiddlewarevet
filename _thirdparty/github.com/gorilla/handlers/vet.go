package main

import (
	"net/http"
	"os"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/gorilla/handlers"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "LoggingHandler",
			Func: func(h http.Handler) http.Handler {
				return handlers.LoggingHandler(os.Stdout, h)
			},
		},
		testing.Middleware{
			Name: "CombinedLoggingHandler",
			Func: func(h http.Handler) http.Handler {
				return handlers.CombinedLoggingHandler(os.Stdout, h)
			},
		},
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
