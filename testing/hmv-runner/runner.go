package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	pkgs, err := findPackages()
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		fmt.Printf("testing %q\n", pkg)
		log, report, ok := runTest(ctx, pkg)
		_ = report
		_ = ok
		os.Stdout.Write(log)
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

func runTest(ctx context.Context, pkg string) ([]byte, []byte, bool) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	var buf bytes.Buffer
	var report bytes.Buffer

	dir, err := ioutil.TempDir("", "test-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	bin := path.Join(dir, "test")

	{ // download dependecies
		// `-d` we are just downloading the missing deps here
		cmd := exec.Command("go", "get", "-d", "-v", "./"+path.Join("_thirdparty", pkg))
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
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
			return buf.Bytes(), nil, false
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
			return buf.Bytes(), nil, false
		}
	}

	{ // run binary
		fmt.Fprintf(&buf, "$ run %s\n", pkg)
		cmd := exec.Command(bin)
		cmd.Stdout = &report
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
			return buf.Bytes(), nil, false
		}
	}

	return buf.Bytes(), report.Bytes(), true
}
