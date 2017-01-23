// +build !go1.8

package http_middleware_vet

import (
	"io"
	"net/http"
	"testing"
)

func vetHandler(t *testing.T, config configData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.Proto != config.Proto {
			t.Errorf("Invalid Proto. (expected %q)", config.Proto)
			t.Logf("VSN: %s (%d/%d) (tls: %v)", r.Proto, r.ProtoMajor, r.ProtoMinor, r.TLS != nil)
		}

		if (r.TLS != nil) != config.TLS {
			t.Errorf("Invalid TLS Info. (expected %v)", config.TLS)
			t.Logf("ALPN: %s (mutual: %v)", r.TLS.NegotiatedProtocol, r.TLS.NegotiatedProtocolIsMutual)
		}

		if _, ok := w.(http.Flusher); ok != config.IsFlusher {
			t.Errorf("Invalid Flusher. (expected %v)", config.IsFlusher)
		}

		if _, ok := w.(http.Hijacker); ok != config.IsHijacker {
			t.Errorf("Invalid Hijacker. (expected %v)", config.IsHijacker)
		}

		if _, ok := w.(http.CloseNotifier); ok != config.IsCloseNotifier {
			t.Errorf("Invalid CloseNotifier. (expected %v)", config.IsCloseNotifier)
		}

		if _, ok := w.(io.ReaderFrom); ok != config.IsReaderFrom {
			t.Errorf("Invalid ReaderFrom. (expected %v)", config.IsReaderFrom)
		}

	}
}
