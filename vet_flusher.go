package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetFlusher(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if _, ok := w.(http.Flusher); ok != config.IsFlusher {
		if config.IsFlusher {
			t.Errorf("ResponseWriter must implement http.Flusher")
		} else {
			t.Errorf("ResponseWriter must not implement http.Flusher")
		}
	}
}
