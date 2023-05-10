// findlinks2 prints the links in an HTML document read from standard input.
package main

import (
	"crawl3/links"
	"fmt"
	"log"
	"os"
)

/*
go build gopl.io/ch8/crawl3

./crawl3 http://gopl.io/
./crawl3 https://golang.org

go run main.go https://golang.org
*/

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide URL(s): `go run main.go http://my-url.com http://my-url2.com ...` ")
	}

	worklist := make(chan []string)  // lists of URLs, may here duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
