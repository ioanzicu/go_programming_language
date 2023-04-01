// The JSON-based web service of the Open Movie Database lets you search
// https://omdbapi.com/ for a movie by name and download its poster image. Write a tool
// poster that downloads the poster image for the movie named on the command line.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Pick the most important fields from JSON response
type Poster struct {
	Title     string
	PosterURL string `json:"Poster"`
	Response  string
	Error     string

	// Year     string
	// Rated    string
	// Released string
	// Runtime  string
	// Genre    string
	// Director string
	// Writer   string
	// Actors   string
	// Plot     string
	// Language string
	// Country  string
	// Awards   string
	// Ratings []Sources{Source, Value}
	// Metascore  string
	// ImdbRating string `json:"imdbRating"`
	// ImdbVotes  string `json:"imdbVotes"`
	// ImdbID     string `json:"imdbID"`
	// Type       string
	// DVD        string
	// BoxOffice  string
	// Production string
	// Website    string
}

// https://omdbapi.com/?t=john+wick&apikey=API_KEY

const apikey = "API_KEY"

func downloadFile(filepath, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	output, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	return err
}

func getImageExtension(posterURL string) string {
	splitedPostURL := strings.Split(posterURL, ".")
	return splitedPostURL[len(splitedPostURL)-1]
}

func getPoster(title string) (*Poster, error) {
	url := fmt.Sprintf("https://omdbapi.com/?t=%s&apikey=%s", title, apikey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("cannot request %s\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("url %s failed: %s", url, resp.Status)
	}

	var result Poster
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("cannot decode: %v", err)
	}

	if result.Response != "True" {
		return nil, fmt.Errorf("Movie: %s Error: %s\n", title, result.Error)
	}

	return &result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the title of movie after main.go like: go run main.go \"john wick\"")
		os.Exit(1)
	}

	title := os.Args[1]
	println(title)

	titleQuery := url.QueryEscape(title)
	poster, err := getPoster(titleQuery)
	if err != nil {
		fmt.Printf("cannot get poster for movie %s, %s", title, err)
		os.Exit(1)
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot get current working dir", err)
		os.Exit(1)
	}

	filepath := path + "-" + titleQuery + "." + getImageExtension(poster.PosterURL)
	err = downloadFile(filepath, poster.PosterURL)
	if err != nil {
		fmt.Printf("cannot download file from %s, %s", poster.PosterURL, err)
		os.Exit(1)
	}
}
