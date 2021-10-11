package httpmiddlewarevet

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

type testCase struct {
	Name string

	EnableForHTTP  bool
	EnableForHTTPS bool
	EnableForHTTP2 bool

	Func testFunc
}

// Report contains test results
type Report struct {
	Name    string
	Proto   string
	Failed  bool
	Skipped bool
	Message string
}

type testReporter struct {
	*Report
}

func (r *testReporter) Proto() string {
	return r.Report.Proto
}

func (r *testReporter) Skip() {
	r.Report.Skipped = true
}

func (r *testReporter) Fail() {
	r.Report.Failed = true
}

func (r *testReporter) FailWithMessage(msg string) {
	r.Report.Failed = true
	r.Report.Message = msg
}

type test interface {
	Proto() string

	Skip()
	Fail()
	FailWithMessage(msg string)
}

type testFunc func(t test, client *http.Client, start func(h http.Handler) string)
type wrapperFunc func(h http.Handler) http.Handler
type newServerFunc func(h http.Handler) *httptest.Server

func newHTTPServer(h http.Handler) *httptest.Server {
	server := httptest.NewUnstartedServer(h)
	server.Start()
	return server
}

func newHTTPSServer(h http.Handler) *httptest.Server {
	server := httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"http/1.1"}}
	server.StartTLS()
	return server
}

func newHTTP2Server(h http.Handler) *httptest.Server {
	server := httptest.NewUnstartedServer(h)
	server.TLS = &tls.Config{NextProtos: []string{"h2", "http/1.1"}}
	server.StartTLS()
	return server
}

func runTest(tc *testCase, tr *testReporter, newServer newServerFunc, wrapper wrapperFunc) {
	var (
		server *httptest.Server
		wg     = &sync.WaitGroup{}
		done   = make(chan struct{})
		rt     = &http.Transport{}
		client = &http.Client{Transport: rt}
	)

	{ // setup default config
		// fails because there is no server running at that address (but used to setup HTTP/2)
		client.Get("http://127.0.0.1:1/")
		// TODO: go1.5 doesn't initialise TLSClientConfig
		if rt.TLSClientConfig == nil {
			rt.TLSClientConfig = &tls.Config{}
		}
		rt.TLSClientConfig.InsecureSkipVerify = true
	}

	defer func() {
		close(done)
		if server != nil {
			server.Close()
		}
	}()

	serve := func(h http.Handler) string {
		if server != nil {
			panic("can only be called once")
		}

		if wrapper != nil {
			h = wrapper(h)
		}

		waiter := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wg.Add(1)
			defer wg.Done()

			h.ServeHTTP(w, r)
		})

		server = newServer(waiter)

		go func() {
			select {
			case <-time.After(1 * time.Minute):
			case <-done:
			}

			server.CloseClientConnections()
			server.Close()
		}()

		return server.URL
	}

	tc.Func(tr, client, serve)
	wg.Wait()
}
