// Modify fetch to also print the HTTP status code, found in resp.Status.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			fmt.Println("http prexif was added")
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		written, err := io.Copy(os.Stdout, resp.Body) // copying to stdout without using an intermediary large buffer
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not copy response: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Written %v bytes\n", written)
		fmt.Printf("Status Code %v\n", resp.StatusCode)
	}
}
