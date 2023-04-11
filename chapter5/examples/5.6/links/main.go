package main

import (
	"fmt"
	"links/links"
	"os"
)

/*
go build .
./findlinks https://golang.org

go run main.go https://golang.org
*/

func main() {
	for _, url := range os.Args[1:] {
		links, err := links.Extract(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "links: %v\n", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}
