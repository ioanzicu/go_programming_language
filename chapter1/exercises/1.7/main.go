// The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a
// buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

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
	}
}
