package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/romainmenke/pusher/link"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "link.Handler",
			Func: func(h http.Handler) http.Handler {
				return link.Handler(h)
			},
		},
	)
}
