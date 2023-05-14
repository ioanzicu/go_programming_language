package memo2

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
https://golang.org, 1.532450831s, 55418 bytes
https://godoc.org, 1.305611667s, 31073 bytes
https://play.golang.org, 1.592710571s, 26683 bytes
http://gopl.io, 1.943628644s, 4154 bytes
https://golang.org, 1.023µs, 55418 bytes
https://godoc.org, 346ns, 31073 bytes
https://play.golang.org, 332ns, 26683 bytes
http://gopl.io, 265ns, 4154 bytes
--- PASS: Test (6.37s)
=== RUN   TestConcurrent
https://godoc.org, 622.692367ms, 31073 bytes
https://golang.org, 1.332637463s, 55418 bytes
http://gopl.io, 2.149691493s, 4154 bytes
https://godoc.org, 2.149726601s, 31073 bytes
https://golang.org, 2.149788253s, 55418 bytes
http://gopl.io, 2.149804842s, 4154 bytes
https://play.golang.org, 3.072175622s, 26683 bytes
https://play.golang.org, 3.072177537s, 26683 bytes
--- PASS: TestConcurrent (3.07s)
PASS
ok      memo2   9.451s







go test -v -race
=== RUN   Test
https://golang.org, 1.265054148s, 55418 bytes
https://godoc.org, 1.163317416s, 31073 bytes
https://play.golang.org, 1.373127726s, 26683 bytes
http://gopl.io, 1.967776891s, 4154 bytes
https://golang.org, 8.693µs, 55418 bytes
https://godoc.org, 7.362µs, 31073 bytes
https://play.golang.org, 6.152µs, 26683 bytes
http://gopl.io, 6.688µs, 4154 bytes
--- PASS: Test (5.77s)
=== RUN   TestConcurrent
https://golang.org, 638.630159ms, 55418 bytes
https://godoc.org, 1.241397567s, 31073 bytes
https://play.golang.org, 2.320128456s, 26683 bytes
http://gopl.io, 3.067370533s, 4154 bytes
https://golang.org, 3.067076428s, 55418 bytes
https://godoc.org, 3.066807027s, 31073 bytes
http://gopl.io, 3.06619038s, 4154 bytes
https://play.golang.org, 3.06636765s, 26683 bytes
--- PASS: TestConcurrent (3.07s)
PASS
ok      memo2   8.869s
*/
