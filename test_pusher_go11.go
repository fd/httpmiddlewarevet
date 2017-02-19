// +build !go1.8

package httpmiddlewarevet

import "net/http"

var pusherTest = &testCase{
	Name: "http.Pusher",
	Func: func(t test, client *http.Client, serve func(h http.Handler) string) {
		t.Skip()
	},
}
