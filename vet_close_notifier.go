package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetCloseNotifier(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if _, ok := w.(http.CloseNotifier); ok != config.IsCloseNotifier {
		if config.IsCloseNotifier {
			t.Errorf("ResponseWriter must implement http.CloseNotifier")
		} else {
			t.Errorf("ResponseWriter must not implement http.CloseNotifier")
		}
	}
}
