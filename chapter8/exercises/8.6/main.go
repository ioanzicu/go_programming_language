package main

import (
	"findlinks/links"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
go build findlinks

./findlinks -depth=2 http://gopl.io/
./findlinks -depth=2 https://golang.org

go run main.go -depth=2 https://golang.org
*/

var maxDepth int

type Link struct {
	url   string
	depth int
}

func crawl(l Link) []Link {
	fmt.Println(l.url)
	var linksList []Link
	if l.depth < maxDepth {
		depth := l.depth + 1
		list, err := links.Extract(l.url)
		if err != nil {
			log.Print(err)
		}

		for _, url := range list {
			linksList = append(linksList, Link{url, depth})
		}
	}
	return linksList
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide URL(s): `go run main.go http://my-url.com http://my-url2.com ...` ")
	}

	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()

	worklist := make(chan []Link)  // lists of URLs, may here duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		var links []Link
		for _, url := range flag.Args() {
			links = append(links, Link{url, 0})
		}
		worklist <- links
	}()

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
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
	close(unseenLinks)
}
