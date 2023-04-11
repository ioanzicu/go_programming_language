// Modify crawl to make local copies of the pages it finds, creating directories as
// necessary. Don’t make copies of pages that come from a different domain. For example, if the
// original page comes from golang.org, save all files from there, but exclude ones from
// vimeo.com.
package main

import (
	"findlinks3/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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
func breadhFirst(f func(item, host string) []string, worklist []string) {
	seen := make(map[string]bool)

	for _, w := range worklist {
		url, err := url.Parse(w)
		if err != nil {
			continue
		}
		host := url.Host

		subworkList := make([]string, 1)
		subworkList[0] = w

		for len(subworkList) > 0 {
			items := subworkList
			subworkList = nil
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					worklist = append(worklist, f(item, host)...)
				}
			}
		}
	}
}

func crawl(url string, host string) []string {
	fmt.Println(url)
	err := savePage(url, host)
	if err != nil {
		log.Printf("Can't save URL \"%s\": %s", url, &err)
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func savePage(rawUrl, host string) error {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}

	if host != url.Host {
		return nil
	}

	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = url.Path
	}

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	resp, err := http.Get(rawUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)

	// check for delayed write errors
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}

	return err
}
