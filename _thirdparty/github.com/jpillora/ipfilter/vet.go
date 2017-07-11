package main

import (
	"log"
	"net/http"

	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/jpillora/ipfilter"
)

type logger struct{}

func (l *logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func main() {

	testing.Run(
		testing.Middleware{
			Name: "Wrap",
			Func: func(h http.Handler) http.Handler {
				return ipfilter.Wrap(h, ipfilter.Options{
					BlockedCountries: []string{"CN", "RU"},
					Logger:           &logger{},
				})
			},
		},
	)
}
