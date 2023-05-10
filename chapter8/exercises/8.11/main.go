// Following the approach of mirroredQuery in Section 8.4.4, implement a vari-
// ant of fetch that requests several URLs concurrently. As soon as the first response arrives,
// cancel the other requests.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type response struct {
	filename string
	n        int64
	err      error
}

// go run mian.go http://gopl.io/

func main() {
	cancel := make(chan struct{})
	resps := make(chan response)
	for _, url := range os.Args[1:] {
		go func(url string) {
			resp := fetch(url, cancel)
			if resp.err != nil {
				fmt.Printf("Failed fetch %s\n", resp.err)
				return
			}
			close(cancel)
			fmt.Printf("Got %s\n", url)
			resps <- resp
		}(url)
	}
	resp := <-resps
	fmt.Printf("Save file %s, Size, %d\n", resp.filename, resp.n)
}

func fetch(url string, cancel <-chan struct{}) response {
	var f *os.File
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response{"", 0, err}
	}

	// Set channel for cancel request
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return response{"", 0, err}
	}

	defer func() {
		resp.Body.Close()
		// Close file, but prefer error from Copy, if any
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err = os.Create(local)
	if err != nil {
		return response{"", 0, err}
	}
	n, err := io.Copy(f, resp.Body)
	return response{local, n, err}
}
