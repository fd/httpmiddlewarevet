package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/fd/httpmiddlewarevet/reports"

	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	pkgs, err := findPackages()
	if err != nil {
		panic(err)
	}

	var r []*reports.UnversionedPackage

	for _, pkg := range pkgs {
		fmt.Printf("testing %q\n", pkg)
		report := runTest(ctx, pkg)
		r = append(r, report)
	}

	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	commitHash := os.Getenv("TRAVIS_COMMIT")
	if commitHash == "" {
		commitHash = "dev"
	}

	reportFile := "./dist/" + commitHash + "/" + runtime.Version() + ".json"

	err = os.MkdirAll(path.Dir(reportFile), 0777)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(reportFile, data, 0666)
	if err != nil {
		panic(err)
	}
}

func findPackages() ([]string, error) {
	var (
		pkgMap = map[string]bool{}
		pkgs   []string
	)

	err := filepath.Walk("_thirdparty", func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		filename = strings.TrimPrefix(filename, "_thirdparty")
		filename = strings.TrimPrefix(filename, "/")
		if filename == "" {
			return nil
		}

		base := path.Base(filename)

		if strings.HasPrefix(base, "_") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasPrefix(base, ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(base, ".go") {
			return nil
		}

		if strings.HasSuffix(base, "_test.go") {
			return nil
		}

		pkg := path.Dir(filename)

		if pkgMap[pkg] {
			return nil
		}

		pkgMap[pkg] = true
		pkgs = append(pkgs, pkg)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Strings(pkgs)
	return pkgs, nil
}

func runTest(ctx context.Context, pkg string) (report *reports.UnversionedPackage) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	report = &reports.UnversionedPackage{Pkg: pkg}

	var buf bytes.Buffer
	var reportBuf bytes.Buffer

	defer func() {
		var handlers []*reports.Handler

		data := reportBuf.Bytes()
		data = bytes.TrimSpace(data)
		if len(data) > 0 {
			err := json.Unmarshal(data, &handlers)
			if err != nil {
				fmt.Fprintf(&buf, "failed to decode report: %s", err)
			}
		}

		report.Log = buf.String()
		report.Handlers = handlers

		var (
			hasFailed  bool
			hasSkipped bool
			hasPassed  bool
		)

		for _, h := range handlers {
			switch h.Status {
			case "failed":
				hasFailed = true
			case "passed":
				hasPassed = true
			case "skipped":
				hasSkipped = true
			}
		}

		switch {
		case hasFailed:
			report.Status = "failed"
		case hasPassed:
			report.Status = "passed"
		case hasSkipped:
			report.Status = "skipped"
		}
	}()

	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	bin := path.Join(dir, "test")

	{ // download dependecies
		fmt.Fprintf(&buf, "$ go get -d %s\n", pkg)
		// `-d` we are just downloading the missing deps here
		cmd := exec.Command("go", "get", "-d", "-v", "./"+path.Join("_thirdparty", pkg))
		cmd.Stderr = io.MultiWriter(os.Stderr, &buf)
		cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
		go func() {
			<-ctx.Done()
			cmd.Process.Kill()
		}()
		err := cmd.Run()
		if err != nil {
			if ctx.Err() == context.Canceled {
				io.WriteString(&buf, "killed\n")
			}
			if ctx.Err() == context.DeadlineExceeded {
				io.WriteString(&buf, "timeout\n")
			}
			return
		}
	}

	{ // build binary
		fmt.Fprintf(&buf, "$ go build %s\n", pkg)
		// `-i` to install deps in GOPATH
		// `-o` to control where the bin is saved
		cmd := exec.Command("go", "build", "-i", "-o", bin, "./"+path.Join("_thirdparty", pkg))
		go func() {
			<-ctx.Done()
			cmd.Process.Kill()
		}()
		result, err := cmd.CombinedOutput()
		buf.Write(result)
		if err != nil {
			if ctx.Err() == context.Canceled {
				io.WriteString(&buf, "killed\n")
			}
			if ctx.Err() == context.DeadlineExceeded {
				io.WriteString(&buf, "timeout\n")
			}
			return
		}
	}

	{ // run binary
		fmt.Fprintf(&buf, "$ run %s\n", pkg)
		cmd := exec.Command(bin)
		cmd.Stdout = &reportBuf
		cmd.Stderr = &buf
		go func() {
			<-ctx.Done()
			cmd.Process.Kill()
		}()
		err := cmd.Run()
		if err != nil {
			if ctx.Err() == context.Canceled {
				io.WriteString(&buf, "killed\n")
			}
			if ctx.Err() == context.DeadlineExceeded {
				io.WriteString(&buf, "timeout\n")
			}
			return
		}
	}

	return
}
