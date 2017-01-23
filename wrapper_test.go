package http_middleware_vet

import (
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	Vet(t, func(h http.Handler) http.Handler {
		return h
	})
}
