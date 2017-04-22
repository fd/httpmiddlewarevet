package main

import (
	"net/http"
	"time"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/romainmenke/pusher/parser"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Handler",
			Func: func(h http.Handler) http.Handler {
				return parser.Handler(h, parser.WithCache(), parser.CacheDuration(time.Second*10))
			},
		},
	)
}
