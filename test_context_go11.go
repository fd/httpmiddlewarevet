// +build !go1.8

package httpmiddlewarevet

import "net/http"

var contextTest = &testCase{
	Name: "Context",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		// r.Context is present since go1.7 but r.Context.Done() doesn't work
		// reliably.
		t.Skip()
	},
}
