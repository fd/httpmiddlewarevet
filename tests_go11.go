// +build !go1.7

package httpmiddlewarevet

import "testing"

func reportResults(t *testing.T, reports []*Report) {
	for _, report := range reports {
		if report.Failed {
			if report.Message == "" {
				t.Errorf("[%s] %s: %s", report.Proto, report.Name, "test failed")
			} else {
				t.Errorf("[%s] %s: %s", report.Proto, report.Name, report.Message)
			}
		} else if report.Skipped {
			t.Logf("[%s] %s: %s", report.Proto, report.Name, "test skipped")
		} else {
			t.Logf("[%s] %s: %s", report.Proto, report.Name, "test passed")
		}
	}
}
