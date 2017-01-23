package http_middleware_vet

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/net/http2"
)

type configData struct {
	Proto string // HTTP/2.0
	TLS   bool

	IsFlusher       bool
	IsHijacker      bool
	IsCloseNotifier bool
	IsReaderFrom    bool
}

var http11Config = configData{
	Proto:           "HTTP/1.1",
	TLS:             false,
	IsFlusher:       true,
	IsHijacker:      true,
	IsCloseNotifier: true,
	IsReaderFrom:    true,
}

var https11Config = configData{
	Proto:           "HTTP/1.1",
	TLS:             true,
	IsFlusher:       true,
	IsHijacker:      true,
	IsCloseNotifier: true,
	IsReaderFrom:    true,
}

var http20Config = configData{
	Proto:           "HTTP/2.0",
	TLS:             true,
	IsFlusher:       true,
	IsHijacker:      false,
	IsCloseNotifier: true,
	IsReaderFrom:    false,
}

// Vet will verify that you correctly implement you http.ResponseWriter wrapper
func Vet(t *testing.T, f func(h http.Handler) http.Handler) {
	t.Run("HTTP/1.1", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, http11Config)))

		server.Start()
		defer server.Close()

		http.Get(server.URL)
	})

	t.Run("TLS/HTTP/1.1", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, https11Config)))

		server.StartTLS()
		defer server.Close()

		client := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}

		client.Get(server.URL)
	})

	t.Run("TLS/HTTP/2.0", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, http20Config)))
		server.TLS = &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		}

		server.StartTLS()
		defer server.Close()

		client := http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
		http2.ConfigureTransport(client.Transport.(*http.Transport))

		client.Get(server.URL)
	})
}

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
