package memo

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
	Get(key string, done <-chan bool) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
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
			value, err := m.Get(url, nil)
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
	defer m.Close()
	Sequential(t, m)
}

// Not concurrency-safe
func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	key := "https://golang.org"
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		v, err := m.Get(key, nil)
		wg1.Done()
		if v == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, v, nil)
		}
	}()
	wg1.Wait()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		done := make(chan bool)
		close(done)
		v, err := m.Get(key, done)
		if v != nil || err == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, "cancled")
		}
		wg2.Done()
	}()
	wg2.Wait()
}

/*
go test -v
=== RUN   Test
https://golang.org, 1.785684562s, 55418 bytes
https://godoc.org, 1.608754604s, 31073 bytes
https://play.golang.org, 1.769099613s, 26683 bytes
http://gopl.io, 2.046621567s, 4154 bytes
https://golang.org, 8.953µs, 55418 bytes
https://godoc.org, 12.605µs, 31073 bytes
https://play.golang.org, 4.292µs, 26683 bytes
http://gopl.io, 3.917µs, 4154 bytes
--- PASS: Test (7.21s)
=== RUN   TestConcurrent
https://godoc.org, 649.400389ms, 31073 bytes
https://godoc.org, 649.365189ms, 31073 bytes
https://golang.org, 668.747763ms, 55418 bytes
https://golang.org, 668.625371ms, 55418 bytes
http://gopl.io, 743.589193ms, 4154 bytes
http://gopl.io, 743.278489ms, 4154 bytes
https://play.golang.org, 967.221264ms, 26683 bytes
https://play.golang.org, 966.909788ms, 26683 bytes
--- PASS: TestConcurrent (0.97s)
=== RUN   TestCancel
--- PASS: TestCancel (0.57s)
PASS
ok      memo    8.751s




go test -v -race
=== RUN   Test
https://golang.org, 1.090410072s, 55418 bytes
https://godoc.org, 948.246024ms, 31073 bytes
https://play.golang.org, 1.418445342s, 26683 bytes
http://gopl.io, 1.658030587s, 4154 bytes
https://golang.org, 363.591µs, 55418 bytes
https://godoc.org, 296.333µs, 31073 bytes
https://play.golang.org, 264.513µs, 26683 bytes
http://gopl.io, 478.318µs, 4154 bytes
--- PASS: Test (5.12s)
=== RUN   TestConcurrent
https://godoc.org, 605.172088ms, 31073 bytes
https://godoc.org, 602.244539ms, 31073 bytes
https://golang.org, 608.091652ms, 55418 bytes
https://golang.org, 610.756388ms, 55418 bytes
http://gopl.io, 831.836946ms, 4154 bytes
http://gopl.io, 835.489578ms, 4154 bytes
https://play.golang.org, 1.199479631s, 26683 bytes
https://play.golang.org, 1.202214561s, 26683 bytes
--- PASS: TestConcurrent (1.20s)
=== RUN   TestCancel
--- PASS: TestCancel (0.72s)
PASS
ok      memo    7.067s
*/
