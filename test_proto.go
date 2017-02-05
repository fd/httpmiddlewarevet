package httpmiddlewarevet

import (
	"fmt"
	"net/http"
)

var protoTest = &testCase{
	Name: "Proto",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if t.Proto() == protoHTTP11TLS {
				if protoHTTP11 != r.Proto {
					t.FailWithMessage(fmt.Sprintf("unexpected request.Proto: %q", r.Proto))
				}
				return
			}

			if t.Proto() != r.Proto {
				t.FailWithMessage(fmt.Sprintf("unexpected request.Proto: %q", r.Proto))
				return
			}

		}))

		client.Get(url)
	},
}
