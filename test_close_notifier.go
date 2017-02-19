package httpmiddlewarevet

import (
	"net/http"
	"time"
)

var closeNotifierTest = &testCase{
	Name: "http.CloseNotifier",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			flusher, ok := w.(http.Flusher)
			if !ok {
				t.FailWithMessage("skipped: http.ResponseWriter is not http.Flusher")
				return
			}

			notifier, ok := w.(http.CloseNotifier)
			if !ok {
				t.FailWithMessage("http.ResponseWriter must implement http.CloseNotifier")
				return
			}

			closed := notifier.CloseNotify()
			deadline := time.After(1 * time.Second)

			w.WriteHeader(200)
			flusher.Flush()
			for {
				select {
				case <-closed:
					return
				case <-deadline:
					t.FailWithMessage("CloseNotify() not propagated")
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
