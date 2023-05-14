package memo3

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
https://golang.org, 1.26450299s, 55418 bytes
https://godoc.org, 1.433895315s, 31073 bytes
https://play.golang.org, 1.526631751s, 26683 bytes
http://gopl.io, 2.147185089s, 4154 bytes
https://golang.org, 1.392µs, 55418 bytes
https://godoc.org, 365ns, 31073 bytes
https://play.golang.org, 281ns, 26683 bytes
http://gopl.io, 413ns, 4154 bytes
--- PASS: Test (6.37s)
=== RUN   TestConcurrent
https://golang.org, 635.623977ms, 55418 bytes
https://godoc.org, 703.892899ms, 31073 bytes
https://golang.org, 713.289777ms, 55418 bytes
https://godoc.org, 793.022764ms, 31073 bytes
https://play.golang.org, 804.160643ms, 26683 bytes
http://gopl.io, 818.044636ms, 4154 bytes
http://gopl.io, 1.005178472s, 4154 bytes
https://play.golang.org, 1.127199703s, 26683 bytes
--- PASS: TestConcurrent (1.13s)
PASS
ok      memo3   7.506s





go test -v -race
=== RUN   Test
https://golang.org, 1.206661897s, 55418 bytes
https://godoc.org, 1.024414115s, 31073 bytes
https://play.golang.org, 1.317699907s, 26683 bytes
http://gopl.io, 1.541900361s, 4154 bytes
https://golang.org, 6.233µs, 55418 bytes
https://godoc.org, 13.721µs, 31073 bytes
https://play.golang.org, 8.588µs, 26683 bytes
http://gopl.io, 6.49µs, 4154 bytes
--- PASS: Test (5.09s)
=== RUN   TestConcurrent
https://godoc.org, 598.935575ms, 31073 bytes
https://golang.org, 616.809452ms, 55418 bytes
https://godoc.org, 692.636134ms, 31073 bytes
http://gopl.io, 798.474211ms, 4154 bytes
https://golang.org, 833.895687ms, 55418 bytes
http://gopl.io, 1.004912103s, 4154 bytes
https://play.golang.org, 1.217437783s, 26683 bytes
https://play.golang.org, 1.220611224s, 26683 bytes
--- PASS: TestConcurrent (1.22s)
PASS
ok      memo3   6.353s
*/
