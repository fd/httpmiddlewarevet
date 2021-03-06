package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Handler",
			Func: func(h http.Handler) http.Handler {
				return h
			},
		},
	)
}
