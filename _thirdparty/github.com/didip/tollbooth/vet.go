package main

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/fd/httpmiddlewarevet/testing"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "LimitHandler",
			Func: func(h http.Handler) http.Handler {
				return tollbooth.LimitHandler(tollbooth.NewLimiter(1, time.Second), h)
			},
		},
	)
}
