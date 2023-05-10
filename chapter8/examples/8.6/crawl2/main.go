// findlinks2 prints the links in an HTML document read from standard input.
package main

import (
	"crawl2/links"
	"fmt"
	"log"
	"os"
)

/*
go build gopl.io/ch8/crawl2

./crawl2 http://gopl.io/
./crawl2 https://golang.org

go run main.go https://golang.org
*/

// tokens is a counting semaphore used to enforce
// a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide URL(s): `go run main.go http://my-url.com http://my-url2.com ...` ")
	}

	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
