package testing

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"

	vet "github.com/fd/httpmiddlewarevet"
	"github.com/fd/httpmiddlewarevet/reports"
)

// Middleware holds a middleware handler to be tested
type Middleware struct {
	Name string
	Func func(h http.Handler) http.Handler
}

// Run all middleware tests
func Run(middlewares ...Middleware) {
	reports := runAll(middlewares...)

	data, err := json.MarshalIndent(reports, "", "  ")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
}

func runAll(middlewares ...Middleware) []*reports.Handler {
	var r []*reports.Handler

	for _, middleware := range middlewares {
		handlerReport := runSingle(middleware)
		r = append(r, handlerReport)
	}

	return r
}

func runSingle(middleware Middleware) *reports.Handler {
	r := vet.Run(middleware.Func)

	handlerReport := &reports.Handler{
		Handler: middleware.Name,
	}

	var (
		handlerHasFailed  bool
		handlerHasSkipped bool
		handlerHasPassed  bool
	)

	for _, group := range groupReports(r, byName) {
		testReport := &reports.Test{Name: group.Key}
		handlerReport.Tests = append(handlerReport.Tests, testReport)

		var (
			testHasFailed  bool
			testHasSkipped bool
			testHasPassed  bool
		)

		for _, report := range group.Reports {
			r := &reports.Proto{
				Name: report.Proto,
			}

			if report.Failed {
				r.Status = "failed"
				r.Message = report.Message
				testHasFailed = true
				handlerHasFailed = true
			} else if report.Skipped {
				r.Status = "skipped"
				r.Message = report.Message
				testHasSkipped = true
				handlerHasSkipped = true
			} else {
				r.Status = "passed"
				testHasPassed = true
				handlerHasPassed = true
			}

			testReport.Protos = append(testReport.Protos, r)
		}

		switch {
		case testHasFailed:
			testReport.Status = "failed"
		case testHasPassed:
			testReport.Status = "passed"
		case testHasSkipped:
			testReport.Status = "skipped"
		}
	}

	switch {
	case handlerHasFailed:
		handlerReport.Status = "failed"
	case handlerHasPassed:
		handlerReport.Status = "passed"
	case handlerHasSkipped:
		handlerReport.Status = "skipped"
	}

	return handlerReport
}

func byName(r *vet.Report) string  { return r.Name }
func byProto(r *vet.Report) string { return r.Proto }

type group struct {
	Key     string
	Reports []*vet.Report
}

func groupReports(r []*vet.Report, grouper func(*vet.Report) string) []*group {
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
