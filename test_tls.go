package httpmiddlewarevet

import "net/http"

var tlsTest = &testCase{
	Name: "TLS",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if t.Proto() == protoHTTP11 {
				if r.TLS != nil {
					t.FailWithMessage("expected request.TLS to be nil")
				}
			}

			if t.Proto() != protoHTTP11 {
				if r.TLS == nil {
					t.FailWithMessage("expected request.TLS to be non-nil")
				}
			}

		}))

		client.Get(url)
	},
}
