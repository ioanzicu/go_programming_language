package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

/*
go run main.go http://gopl.io
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide url argument")
		os.Exit(1)
	}

	filename, n, err := fetch(os.Args[1])
	if err != nil {
		fmt.Printf("err: %s", err)
		os.Exit(1)
	}

	fmt.Println("Filename:", filename)
	fmt.Println("Copied bytes:", n)
}

// Fetch downloads the URL and returns the name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	// the f.Close will be called after the function
	// updates the named return values
	// if error, it will update the returned error value
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
