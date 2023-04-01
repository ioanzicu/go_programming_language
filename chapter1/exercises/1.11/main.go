// Try fetchall with longer argument lists, such as samples from the top million
// web sites available at alexa.com. How does the program behave if a web site just doesn’t
// respond? (Section 8.9 describes mechanisms for coping in such cases.)
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*
fetches many URLs, all concurrently,
so that the process will take no longer than the
longest fetch rather than the sum of all the fetch times.
*/

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*

go run main.go https://golang.org http://gopl.io https://godoc.org https://facebook.com https://fonts.googleapis.com https://googletagmanager.com https://s.w.org https://fonts.gstatic.com

Get "https://s.w.org": dial tcp 192.0.77.48:443: i/o timeout
Get "https://godoc.org": dial tcp 34.160.232.96:443: i/o timeout
Get "https://googletagmanager.com": dial tcp 142.251.1.97:443: i/o timeout
Get "https://facebook.com": dial tcp 31.13.72.36:443: i/o timeout
Get "https://fonts.googleapis.com": dial tcp 74.125.205.95:443: i/o timeout
Get "https://fonts.gstatic.com": dial tcp 216.58.210.163:443: i/o timeout
Get "http://gopl.io": dial tcp 52.88.110.197:80: i/o timeout
Get "https://golang.org": dial tcp 216.58.209.209:443: i/o timeout
30.03s elapsed

*/
