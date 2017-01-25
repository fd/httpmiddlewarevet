// +build !go1.7

package httpmiddlewarevet

import "net/http"

var contextTest = &testCase{
	Name: "Context",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		t.Skip()
	},
}
