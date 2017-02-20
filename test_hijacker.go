package httpmiddlewarevet

import "net/http"

var hijackerTest = &testCase{
	Name: "http.Hijacker",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, ok := w.(http.Hijacker)
			if t.Proto() == protoHTTP20 && ok {
				t.FailWithMessage("http.ResponseWriter must not implement http.Hijacker")
			}
			if t.Proto() != protoHTTP20 && !ok {
				t.FailWithMessage("http.ResponseWriter must implement http.Hijacker")
			}

		}))

		client.Get(url)
	},
}
