package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetProto(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	if r.Proto != config.Proto {
		t.Errorf("Invalid Proto. (expected %q)", config.Proto)
		t.Logf("VSN: %s (%d/%d) (tls: %v)", r.Proto, r.ProtoMajor, r.ProtoMinor, r.TLS != nil)
	}
}
