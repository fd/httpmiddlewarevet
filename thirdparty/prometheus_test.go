package thirdparty

import (
	"net/http"
	"testing"

	"github.com/fd/httpmiddlewarevet"
	"github.com/prometheus/client_golang/prometheus"
)

func TestPrometheus(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return prometheus.InstrumentHandler("vet", h)
	})
}
