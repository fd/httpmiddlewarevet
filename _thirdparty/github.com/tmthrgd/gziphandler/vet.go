package main

import (
	"github.com/fd/httpmiddlewarevet/testing"
	"github.com/tmthrgd/gziphandler"
)

func main() {
	testing.Run(
		testing.Middleware{
			Name: "GzipHandler",
			Func: gziphandler.GzipHandler,
		},
	)
}
