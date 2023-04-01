// Find a web site that produces a large amount of data. Investigate caching by
// running fetchall twice in succession to see whether the reported time changes much. Do
// you get the same content each time? Modify fetchall to print its output to a file so it can be
// examined.
package main

import (
	"fmt"
	"io"
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

	w, err := os.Create(url[8:14] + ".txt")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	nbytes, err := io.Copy(w, resp.Body)

	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*

go run chapter1/examples/1.10/fetchall/main.go https://golang.org http://gopl.io https://godoc.org

1.60s   30531 https://godoc.org
1.87s   55220 https://golang.org
1.92s    4154 http://gopl.io
1.93s elapsed

*/
