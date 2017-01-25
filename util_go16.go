// +build go1.6

package httpmiddlewarevet

func runTests(tc *testCase, wrapper wrapperFunc) []*Report {
	var reports = make([]*Report, 0, 3)

	if tc.EnableForHTTP == false &&
		tc.EnableForHTTPS == false &&
		tc.EnableForHTTP2 == false {
		tc.EnableForHTTP = true
		tc.EnableForHTTPS = true
		tc.EnableForHTTP2 = true
	}

	r0 := &Report{Name: tc.Name, Proto: "HTTP/1.1"}
	reports = append(reports, r0)
	if tc.EnableForHTTP {
		runTest(tc, &testReporter{r0}, newHTTPServer, wrapper)
	} else {
		r0.Message = "test disabled"
	}

	r1 := &Report{Name: tc.Name, Proto: "HTTP/1.1+TLS"}
	reports = append(reports, r1)
	if tc.EnableForHTTPS {
		runTest(tc, &testReporter{r1}, newHTTPSServer, wrapper)
	} else {
		r1.Message = "test disabled"
	}

	r2 := &Report{Name: tc.Name, Proto: "HTTP/2.0"}
	reports = append(reports, r2)
	if tc.EnableForHTTP2 {
		runTest(tc, &testReporter{r2}, newHTTP2Server, wrapper)
	} else {
		r2.Message = "test disabled"
	}

	return reports
}
