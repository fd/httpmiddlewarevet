package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/justinas/alice"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Then",
			Func: func(h http.Handler) http.Handler {
				return alice.New(func(h http.Handler) http.Handler { return h }).Then(h)
			},
		},
	)
}
