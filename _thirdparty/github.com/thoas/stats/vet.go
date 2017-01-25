package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/thoas/stats"
)

func main() {
	testing.Run(
		"github.com/thoas/stats",

		testing.Middleware{
			Name: "Handler",
			Func: func(h http.Handler) http.Handler {
				return stats.New().Handler(h)
			},
		},
	)
}
