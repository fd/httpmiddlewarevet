package httpmiddlewarevet

import "net/http"

var flusherTest = &testCase{
	Name: "Flusher",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, ok := w.(http.Flusher)
			if !ok {
				t.FailWithMessage("http.ResponseWriter must implement http.Flusher")
				return
			}

		}))

		client.Get(url)
	},
}
