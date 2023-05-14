package memo5

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

// ---
// Package memotest provides common functions for
// testing various designs of the memo package.

// !+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
}

// ---

func Test(t *testing.T) {
	m := New(httpGetBody)
	Sequential(t, m)
}

// Not concurrency-safe
func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	Concurrent(t, m)
}

/*
go test -v
=== RUN   Test
https://golang.org, 1.359058977s, 55418 bytes
https://godoc.org, 1.354408301s, 31073 bytes
https://play.golang.org, 1.613975095s, 26683 bytes
http://gopl.io, 2.579590466s, 4154 bytes
https://golang.org, 24.601µs, 55418 bytes
https://godoc.org, 10.619µs, 31073 bytes
https://play.golang.org, 10.195µs, 26683 bytes
http://gopl.io, 9.955µs, 4154 bytes
--- PASS: Test (6.91s)
=== RUN   TestConcurrent
https://godoc.org, 537.248214ms, 31073 bytes
https://godoc.org, 537.157442ms, 31073 bytes
https://golang.org, 571.158754ms, 55418 bytes
https://golang.org, 570.83155ms, 55418 bytes
http://gopl.io, 605.996172ms, 4154 bytes
http://gopl.io, 606.820542ms, 4154 bytes
https://play.golang.org, 1.001635679s, 26683 bytes
https://play.golang.org, 1.002479522s, 26683 bytes
--- PASS: TestConcurrent (1.00s)
PASS
ok      memo5   7.913s



go test -v -race
=== RUN   Test
https://golang.org, 1.250721048s, 55418 bytes
https://godoc.org, 1.098648773s, 31073 bytes
https://play.golang.org, 1.555376267s, 26683 bytes
http://gopl.io, 1.808558831s, 4154 bytes
https://golang.org, 379.646µs, 55418 bytes
https://godoc.org, 307.442µs, 31073 bytes
https://play.golang.org, 274.627µs, 26683 bytes
http://gopl.io, 305.891µs, 4154 bytes
--- PASS: Test (5.72s)
=== RUN   TestConcurrent
https://godoc.org, 615.352619ms, 31073 bytes
https://godoc.org, 617.543381ms, 31073 bytes
https://golang.org, 649.500078ms, 55418 bytes
https://golang.org, 651.114923ms, 55418 bytes
http://gopl.io, 754.987415ms, 4154 bytes
http://gopl.io, 752.319857ms, 4154 bytes
https://play.golang.org, 1.152011164s, 26683 bytes
https://play.golang.org, 1.154620612s, 26683 bytes
--- PASS: TestConcurrent (1.16s)
PASS
ok      memo5   6.909s
*/
