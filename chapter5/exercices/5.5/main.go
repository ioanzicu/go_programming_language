// Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

/*
go build .
./findlinks1 https://golang.org

go run main.go https://golang.org
*/

func main() {
	for _, url := range os.Args[1:] {
		imgCounter, wordCounter, err := countWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
			continue
		}

		fmt.Println("Images found: ", imgCounter)

		fmt.Printf("\n\ncount\tword")
		for word, count := range wordCounter {
			fmt.Printf("%d\t%q\n", count, word)
		}
	}
}

// countWordsAndImages performs an HTTP GET request for url, parses the
// response as HTML, and count images and words.
func countWordsAndImages(url string) (int, map[string]uint, error) {
	imgCounter := 0
	wordCounter := map[string]uint{}

	resp, err := http.Get(url)
	if err != nil {
		return imgCounter, wordCounter, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return imgCounter, wordCounter, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return imgCounter, wordCounter, fmt.Errorf("parsing %s as HTML: %s", url, err)
	}

	count(wordCounter, &imgCounter, doc)

	return imgCounter, wordCounter, nil
}

func count(wordCounter map[string]uint, imgCounter *int, n *html.Node) {
	// count images
	if n.Type == html.ElementNode && n.Data == "img" {
		*imgCounter++
	}

	// count words
	if n.Type == html.TextNode {

		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)

		for input.Scan() {
			wordCounter[input.Text()]++
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(wordCounter, imgCounter, c)
	}
}
