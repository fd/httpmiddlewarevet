package httpmiddlewarevet

import "net/http"

type stringWriter interface {
	WriteString(s string) (n int, err error)
}

var writeStringTest = &testCase{
	Name: "WriteString()",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		url := serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, ok := w.(stringWriter)
			if t.Proto() == protoHTTP20 && !ok {
				t.FailWithMessage("http.ResponseWriter must implement WriteString(s string) (n int, err error)")
			}
			if t.Proto() != protoHTTP20 && !ok {
				t.FailWithMessage("http.ResponseWriter must implement WriteString(s string) (n int, err error)")
			}

		}))

		client.Get(url)
	},
}
