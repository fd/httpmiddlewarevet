package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
)

func main() {
	testing.Run(
		"net/http",

		testing.Middleware{
			Name: "Handler",
			Func: func(h http.Handler) http.Handler {
				return h
			},
		},
	)
}
