package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetHijacker(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if _, ok := w.(http.Hijacker); ok != config.IsHijacker {
		if config.IsHijacker {
			t.Errorf("ResponseWriter must implement http.Hijacker")
		} else {
			t.Errorf("ResponseWriter must not implement http.Hijacker")
		}
	}
}
