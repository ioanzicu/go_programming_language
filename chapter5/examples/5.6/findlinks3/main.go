// findlinks2 prints the links in an HTML document read from standard input.
package main

import (
	"findlinks3/links"
	"fmt"
	"log"
	"os"
)

/*
go build .
./findlinks3 https://golang.org

go run main.go https://golang.org
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide URL(s): `go run main.go http://my-url.com http://my-url2.com ...` ")
	}
	breadhFirst(crawl, os.Args[1:])
}

// breadhFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadhFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
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
