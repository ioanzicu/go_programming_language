package main

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

var HTTPGetBody = httpGetBody

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

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

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
https://golang.org, 1.26455638s, 55418 bytes
https://godoc.org, 1.205901159s, 31073 bytes
https://play.golang.org, 1.353780885s, 26683 bytes
http://gopl.io, 1.878171016s, 4154 bytes
https://golang.org, 1.431µs, 55418 bytes
https://godoc.org, 346ns, 31073 bytes
https://play.golang.org, 312ns, 26683 bytes
http://gopl.io, 293ns, 4154 bytes
--- PASS: Test (5.70s)
=== RUN   TestConcurrent
https://godoc.org, 585.873058ms, 31073 bytes
https://golang.org, 642.985362ms, 55418 bytes
https://golang.org, 684.491794ms, 55418 bytes
https://godoc.org, 702.737741ms, 31073 bytes
http://gopl.io, 713.903856ms, 4154 bytes
http://gopl.io, 1.047070668s, 4154 bytes
https://play.golang.org, 1.118458086s, 26683 bytes
https://play.golang.org, 1.129629777s, 26683 bytes
--- PASS: TestConcurrent (1.13s)
PASS
ok      memo1   6.839s










go test -race -v
=== RUN   Test
https://golang.org, 1.534339745s, 55418 bytes
https://godoc.org, 1.586810538s, 31073 bytes
https://play.golang.org, 1.47604199s, 26683 bytes
http://gopl.io, 1.731622797s, 4154 bytes
https://golang.org, 3.226µs, 55418 bytes
https://godoc.org, 1.565µs, 31073 bytes
https://play.golang.org, 1.473µs, 26683 bytes
http://gopl.io, 4.917µs, 4154 bytes
--- PASS: Test (6.33s)
=== RUN   TestConcurrent
https://godoc.org, 675.966063ms, 31073 bytes
https://golang.org, 711.780309ms, 55418 bytes
https://godoc.org, 765.879158ms, 31073 bytes
==================
WARNING: DATA RACE
Write at 0x00c0001b1290 by goroutine 91:
  runtime.mapassign_faststr()
      /usr/lib/go-1.18/src/runtime/map_faststr.go:203 +0x0
  memo1.(*Memo).Get()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo.go:28 +0x124
  memo1.Concurrent.func1()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:80 +0xea
  memo1.Concurrent.func2()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:87 +0x58

Previous write at 0x00c0001b1290 by goroutine 94:
  runtime.mapassign_faststr()
      /usr/lib/go-1.18/src/runtime/map_faststr.go:203 +0x0
  memo1.(*Memo).Get()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo.go:28 +0x124
  memo1.Concurrent.func1()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:80 +0xea
  memo1.Concurrent.func2()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:87 +0x58

Goroutine 91 (running) created at:
  memo1.Concurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:77 +0x7d
  memo1.TestConcurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:102 +0xc4
  testing.tRunner()
      /usr/lib/go-1.18/src/testing/testing.go:1439 +0x213
  testing.(*T).Run.func1()
      /usr/lib/go-1.18/src/testing/testing.go:1486 +0x47

Goroutine 94 (finished) created at:
  memo1.Concurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:77 +0x7d
  memo1.TestConcurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:102 +0xc4
  testing.tRunner()
      /usr/lib/go-1.18/src/testing/testing.go:1439 +0x213
  testing.(*T).Run.func1()
      /usr/lib/go-1.18/src/testing/testing.go:1486 +0x47
==================
http://gopl.io, 771.950461ms, 4154 bytes
https://golang.org, 777.647657ms, 55418 bytes
http://gopl.io, 911.573139ms, 4154 bytes
https://play.golang.org, 1.021030811s, 26683 bytes
==================
WARNING: DATA RACE
Write at 0x00c000106de8 by goroutine 96:
  memo1.(*Memo).Get()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo.go:28 +0x135
  memo1.Concurrent.func1()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:80 +0xea
  memo1.Concurrent.func2()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:87 +0x58

Previous write at 0x00c000106de8 by goroutine 89:
  memo1.(*Memo).Get()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo.go:28 +0x135
  memo1.Concurrent.func1()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:80 +0xea
  memo1.Concurrent.func2()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:87 +0x58

Goroutine 96 (running) created at:
  memo1.Concurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:77 +0x7d
  memo1.TestConcurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:102 +0xc4
  testing.tRunner()
      /usr/lib/go-1.18/src/testing/testing.go:1439 +0x213
  testing.(*T).Run.func1()
      /usr/lib/go-1.18/src/testing/testing.go:1486 +0x47

Goroutine 89 (finished) created at:
  memo1.Concurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:77 +0x7d
  memo1.TestConcurrent()
      /home/ioan/Documents/go_programming_language/chapter9/examples/9.7/memo1/memo_test.go:102 +0xc4
  testing.tRunner()
      /usr/lib/go-1.18/src/testing/testing.go:1439 +0x213
  testing.(*T).Run.func1()
      /usr/lib/go-1.18/src/testing/testing.go:1486 +0x47
==================
https://play.golang.org, 1.021745561s, 26683 bytes
    testing.go:1312: race detected during execution of test
--- FAIL: TestConcurrent (1.03s)
=== CONT
    testing.go:1312: race detected during execution of test
FAIL
exit status 1
FAIL    memo1   7.368s
*/
