package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetTLS(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if (r.TLS != nil) != config.TLS {
		t.Errorf("Invalid TLS Info. (expected %v)", config.TLS)
		t.Logf("ALPN: %s (mutual: %v)", r.TLS.NegotiatedProtocol, r.TLS.NegotiatedProtocolIsMutual)
	}
}
