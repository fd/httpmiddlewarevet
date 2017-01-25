package thirdparty

import (
	"net/http"
	"testing"

	"github.com/fd/httpmiddlewarevet"
	"github.com/thoas/stats"
)

func TestThoasStats(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return stats.New().Handler(h)
	})
}
