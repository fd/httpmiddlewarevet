package thirdparty

import (
	"net/http"
	"testing"

	"github.com/NYTimes/gziphandler"
	"github.com/fd/httpmiddlewarevet"
)

func TestMYTimesGzipHandler(t *testing.T) {
	httpmiddlewarevet.Vet(t, func(h http.Handler) http.Handler {
		return gziphandler.GzipHandler(h)
	})
}
