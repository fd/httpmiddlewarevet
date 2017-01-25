package main

import (
	"github.com/NYTimes/gziphandler"
	"github.com/fd/httpmiddlewarevet/testing"
)

func main() {
	testing.Run(
		"github.com/NYTimes/gziphandler",

		testing.Middleware{
			Name: "GzipHandler",
			Func: gziphandler.GzipHandler,
		},
	)
}
