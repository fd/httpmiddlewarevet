package httpmiddlewarevet

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

type configData struct {
	Proto string // HTTP/2.0
	TLS   bool

	IsFlusher       bool
	IsHijacker      bool
	IsCloseNotifier bool
	IsReaderFrom    bool
	IsPusher        bool
}

var http11Config = configData{
	Proto:           "HTTP/1.1",
	TLS:             false,
	IsFlusher:       true,
	IsHijacker:      true,
	IsCloseNotifier: true,
	IsReaderFrom:    true,
	IsPusher:        false,
}

var https11Config = configData{
	Proto:           "HTTP/1.1",
	TLS:             true,
	IsFlusher:       true,
	IsHijacker:      true,
	IsCloseNotifier: true,
	IsReaderFrom:    true,
	IsPusher:        false,
}

var http20Config = configData{
	Proto:           "HTTP/2.0",
	TLS:             true,
	IsFlusher:       true,
	IsHijacker:      false,
	IsCloseNotifier: true,
	IsReaderFrom:    false,
	IsPusher:        true,
}

// Vet will verify that you correctly implement your http.ResponseWriter wrapper.
func Vet(t *testing.T, f func(h http.Handler) http.Handler) {
	t.Run("HTTP/1.1", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, http11Config)))

		server.Start()
		defer server.Close()

		http.Get(server.URL)
	})

	t.Run("TLS/HTTP/1.1", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, https11Config)))
		server.TLS = &tls.Config{
			NextProtos: []string{"http/1.1"},
		}

		server.StartTLS()
		defer server.Close()

		rt := &http.Transport{}
		client := http.Client{Transport: rt}
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		rt.TLSClientConfig.InsecureSkipVerify = true

		client.Get(server.URL)
	})

	t.Run("TLS/HTTP/2.0", func(t *testing.T) {
		server := httptest.NewUnstartedServer(f(vetHandler(t, http20Config)))
		server.TLS = &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		}

		server.StartTLS()
		defer server.Close()

		rt := &http.Transport{}
		client := http.Client{Transport: rt}
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		rt.TLSClientConfig.InsecureSkipVerify = true

		client.Get(server.URL)
	})
}
