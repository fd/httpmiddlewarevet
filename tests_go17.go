// +build go1.7

package httpmiddlewarevet

import "testing"

func reportResults(t *testing.T, reports []*Report) {
	for _, group := range groupReports(reports, byName) {
		t.Run(group.Key, func(t *testing.T) {
			for _, report := range group.Reports {
				t.Run(report.Proto, func(t *testing.T) {
					if report.Failed {
						if report.Message == "" {
							t.Errorf("[%s] %s: %s", report.Proto, report.Name, "test failed")
						} else {
							t.Errorf("[%s] %s: %s", report.Proto, report.Name, report.Message)
						}
					} else if report.Skipped {
						t.Skip()
					}
				})
			}
		})
	}
}
