// +build go1.8

package httpmiddlewarevet

import "net/http"

var pusherTest = &testCase{
	Name: "Pusher",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		var isPusher = map[string]bool{
			"HTTP/1.1":     false,
			"HTTP/1.1+TLS": false,
			"HTTP/2.0":     true,
		}

		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, ok := w.(http.Pusher)

			if isPusher[t.Proto()] && !ok {
				t.FailWithMessage("http.ResponseWriter must implement http.Pusher")
				return
			}

			if !isPusher[t.Proto()] && ok {
				t.FailWithMessage("http.ResponseWriter must not implement http.Pusher")
				return
			}

		}))

		client.Get(url)
	},
}
