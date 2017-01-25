package httpmiddlewarevet

import (
	"io"
	"net/http"
	"testing"
)

func vetReaderFrom(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if _, ok := w.(io.ReaderFrom); ok != config.IsReaderFrom {
		if config.IsReaderFrom {
			t.Errorf("ResponseWriter must implement io.ReaderFrom")
		} else {
			t.Errorf("ResponseWriter must not implement io.ReaderFrom")
		}
	}
}
