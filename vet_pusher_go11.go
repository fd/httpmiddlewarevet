// +build go1.8

package httpmiddlewarevet

import (
	"net/http"
	"testing"
)

func vetPusher(t *testing.T, config configData, w http.ResponseWriter, r *http.Request) {
	// no http.Pusher in pre 1.8
}
