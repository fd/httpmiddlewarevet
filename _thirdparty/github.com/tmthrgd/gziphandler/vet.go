package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/tmthrgd/gziphandler"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Gzip",
			Func: func(h http.Handler) http.Handler {
				return gziphandler.Gzip(h)
			},
		},
	)
}
