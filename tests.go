package httpmiddlewarevet

import (
	"net/http"
	"sort"
	"testing"
)

var tests = []*testCase{
	protoTest,
	tlsTest,
	readerFromTest,
	writeStringTest,
	hijackerTest,
	flusherTest,
	closeNotifierTest,
	contextTest,
	pusherTest,
}

// Run generates reports for a middleware handler
func Run(middleware func(h http.Handler) http.Handler) []*Report {
	var reports []*Report
	for _, test := range tests {
		reports = append(reports, runTests(test, middleware)...)
	}
	return reports
}

// Vet will verify that you correctly implement your http.ResponseWriter wrapper.
func Vet(t *testing.T, middleware func(h http.Handler) http.Handler) {
	reports := Run(middleware)
	reportResults(t, reports)
}

func byName(r *Report) string  { return r.Name }
func byProto(r *Report) string { return r.Proto }

type group struct {
	Key     string
	Reports []*Report
}

func groupReports(r []*Report, grouper func(*Report) string) []*group {
	var (
		keys    []string
		entries = map[string]*group{}
		groups  []*group
	)

	for _, report := range r {
		key := grouper(report)
		if g, ok := entries[key]; ok {
			g.Reports = append(g.Reports, report)
		} else {
			g = &group{Key: key}
			g.Reports = append(g.Reports, report)
			entries[key] = g
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	for _, key := range keys {
		groups = append(groups, entries[key])
	}

	return groups
}
