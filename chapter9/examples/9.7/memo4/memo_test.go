package memo4

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
https://golang.org, 1.753717074s, 55418 bytes
https://godoc.org, 2.389243661s, 31073 bytes
https://play.golang.org, 1.497384256s, 26683 bytes
http://gopl.io, 2.134002355s, 4154 bytes
https://golang.org, 1.706µs, 55418 bytes
https://godoc.org, 957ns, 31073 bytes
https://play.golang.org, 708ns, 26683 bytes
http://gopl.io, 693ns, 4154 bytes
--- PASS: Test (7.77s)
=== RUN   TestConcurrent
https://godoc.org, 563.737788ms, 31073 bytes
https://godoc.org, 563.622362ms, 31073 bytes
https://golang.org, 606.340727ms, 55418 bytes
https://golang.org, 606.227224ms, 55418 bytes
http://gopl.io, 680.233056ms, 4154 bytes
http://gopl.io, 680.191912ms, 4154 bytes
https://play.golang.org, 937.540153ms, 26683 bytes
https://play.golang.org, 937.481561ms, 26683 bytes
--- PASS: TestConcurrent (0.94s)
PASS
ok      memo4   8.716s





 go test -v -race
=== RUN   Test
https://golang.org, 1.310025176s, 55418 bytes
https://godoc.org, 1.614111075s, 31073 bytes
https://play.golang.org, 1.530048648s, 26683 bytes
http://gopl.io, 2.211910322s, 4154 bytes
https://golang.org, 9.663µs, 55418 bytes
https://godoc.org, 5.925µs, 31073 bytes
https://play.golang.org, 6.709µs, 26683 bytes
http://gopl.io, 4.08µs, 4154 bytes
--- PASS: Test (6.67s)
=== RUN   TestConcurrent
https://godoc.org, 562.054948ms, 31073 bytes
https://godoc.org, 558.445547ms, 31073 bytes
https://golang.org, 600.174552ms, 55418 bytes
https://golang.org, 596.712588ms, 55418 bytes
http://gopl.io, 701.232439ms, 4154 bytes
http://gopl.io, 699.154309ms, 4154 bytes
https://play.golang.org, 1.173644245s, 26683 bytes
https://play.golang.org, 1.170906464s, 26683 bytes
--- PASS: TestConcurrent (1.18s)
PASS
ok      memo4   7.880s
*/
