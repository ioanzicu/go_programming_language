// findlinks2 prints the links in an HTML document read from standard input.
package main

import (
	"crawl1/links"
	"fmt"
	"log"
	"os"
)

/*
go build gopl.io/ch8/crawl1

./crawl1 http://gopl.io/
./crawl1 https://golang.org

go run main.go https://golang.org


http://gopl.io/
https://golang.org/help/
https://golang.org/doc/
https://golang.org/blog/
...
2015/07/15 18:22:12 Get ...: dial tcp: lookup blog.golang.org: no such host
2015/07/15 18:22:12 Get ...: dial tcp 23.21.222.120:443: socket:
too many open files
...

*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide URL(s): `go run main.go http://my-url.com http://my-url2.com ...` ")
	}

	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
