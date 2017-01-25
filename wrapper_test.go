package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	Vet(t, func(h http.Handler) http.Handler {
		return h
	})
}
