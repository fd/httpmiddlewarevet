// +build go1.8

package httpmiddlewarevet

import (
	"net/http"
	"time"
)

var contextTest = &testCase{
	Name: "context.Context",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			flusher, ok := w.(http.Flusher)
			if !ok {
				t.FailWithMessage("skipped: http.ResponseWriter is not http.Flusher")
				return
			}

			deadline := time.After(1 * time.Second)

			w.WriteHeader(200)
			flusher.Flush()
			for {
				select {
				case <-r.Context().Done():
					return
				case <-deadline:
					t.FailWithMessage("Context().Done() not propagated")
					return
				}
			}
		}))

		res, err := client.Get(url)
		if err != nil {
			panic(err)
		}

		time.Sleep(200 * time.Millisecond)
		res.Body.Close()
	},
}
