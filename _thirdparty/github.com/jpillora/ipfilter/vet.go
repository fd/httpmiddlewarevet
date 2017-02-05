package main

import (
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/jpillora/ipfilter"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "Wrap",
			Func: func(h http.Handler) http.Handler {
				return ipfilter.Wrap(h, ipfilter.Options{
					BlockedCountries: []string{"CN", "RU"},
				})
			},
		},
	)
}
