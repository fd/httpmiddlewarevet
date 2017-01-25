package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetPusher(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if _, ok := w.(http.Pusher); ok != config.IsPusher {
		if config.IsPusher {
			t.Errorf("ResponseWriter must implement http.Pusher")
		} else {
			t.Errorf("ResponseWriter must not implement http.Pusher")
		}
	}
}
