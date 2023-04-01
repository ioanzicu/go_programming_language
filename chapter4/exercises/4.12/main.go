// The popular web comic xkcd has a JSON interface. For example, a request to
// https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
// many favorites. Download each URL (once!) and build an offline index. Write a tool xkcd
// that, using this index, prints the URL and transcript of each comic that matches a search term
// provided on the command line.
package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Comics struct {
	Month      string
	Number     int `json:"num"`
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	ImageURL   string `json:"img"`
	Title      string
	Day        string
}

// buildsURLs - returns a slice of xkcd URLs from 1 to number inclusive
func buildURLs(number int) []string {
	urls := []string{}
	for i := 1; i <= number; i++ {
		url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", i)
		urls = append(urls, url)
	}
	return urls
}

// getComics - query comics xkcd form given url
func getComics(url string) (*Comics, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("url %s failed: %s", url, resp.Status)
	}

	var result Comics
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// fileExists - check if file with such name exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// pullComics - pull comics from remote URLs and store them in local .gob file
func pullComics(fileName string, comicsCount int) error {
	comicsList := map[int]*Comics{}
	URLs := buildURLs(comicsCount)

	for _, url := range URLs {
		fmt.Printf("Load URL %s\n", url)
		result, err := getComics(url)
		if err != nil {
			return err
		}
		comicsList[result.Number] = result
	}

	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %s", fileName, err)
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(comicsList)
	if err != nil {
		return fmt.Errorf("cannot encode %s", err)
	}

	return nil
}

// loadComicsFromLocal - loads the comics from local .gob file
func loadComicsFromLocal(fileName string) (map[int]*Comics, error) {
	f1, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("cannot open file %s: %s", fileName, err)
		return nil, err
	}
	decoder := gob.NewDecoder(f1)

	comicsMap := map[int]*Comics{}
	err = decoder.Decode(&comicsMap)
	if err != nil {
		fmt.Printf("cannot decode %s", err)
		return nil, err
	}

	return comicsMap, nil
}

// printCompics - prints the given Comics struct fieds
func printCompics(comics *Comics) {
	fmt.Printf("\nTitle %s\n", comics.Title)
	fmt.Printf("Transcript %s\n", comics.Transcript)
	fmt.Printf("Month %s\n", comics.Month)
	fmt.Printf("Number %d\n", comics.Number)
	fmt.Printf("Link %s\n", comics.Link)
	fmt.Printf("Year %s\n", comics.Year)
	fmt.Printf("News %s\n", comics.News)
	fmt.Printf("SafeTitle %s\n", comics.SafeTitle)
	fmt.Printf("Alt %s\n", comics.Alt)
	fmt.Printf("ImageURL %s\n", comics.ImageURL)
	fmt.Printf("Day %s\n", comics.Day)
}

// go run chapter4/exercises/4.12/main.go
func main() {

	// assume that if there is no such file, then pull the content and create the file
	fileName := "xkcdLocal.gob"
	if !fileExists(fileName) {
		fmt.Println("Pull comics from the internet")

		numberOfComics := 100

		err := pullComics(fileName, numberOfComics)
		if err != nil {
			log.Fatalf("cannot pull comics: %s", err)
			os.Exit(1)
		}
	}

	// back-up json data
	comicsMap, err := loadComicsFromLocal(fileName)
	if err != nil {
		log.Fatalf("cannot get comics from local: %s", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please enter the comics number between 1 to 100")
		os.Exit(1)
	}

	number := os.Args[1]
	comicsNumber, err := strconv.ParseInt(number, 10, 32)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Printf("\n\nComics with number %d\n", comicsNumber)
	comics, ok := comicsMap[int(comicsNumber)]
	if !ok {
		log.Fatalf("cannot find comics with number %d", comicsNumber)
		os.Exit(1)
	}

	printCompics(comics)
}
