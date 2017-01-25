package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "InstrumentHandler",
			Func: func(h http.Handler) http.Handler {
				return prometheus.InstrumentHandler("vet", h)
			},
		},
	)
}
