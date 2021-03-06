package httpmiddlewarevet

import (
	"io"
	"net/http"
)

var readerFromTest = &testCase{
	Name: "io.ReaderFrom",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, ok := w.(io.ReaderFrom)
			if t.Proto() == protoHTTP20 && ok {
				t.FailWithMessage("http.ResponseWriter must not implement io.ReaderFrom")
			}
			if t.Proto() != protoHTTP20 && !ok {
				t.FailWithMessage("http.ResponseWriter must implement io.ReaderFrom")
			}

		}))

		client.Get(url)
	},
}
