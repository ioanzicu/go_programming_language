package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Casablanca",
		Year:   1942,
		Color:  false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"}},
	{
		Title:  "Bullitt",
		Year:   1968,
		Color:  true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
	},
	// ...
}

func main() {
	// Go struct -> JSON
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n\n", data)

	var titles []struct{ Title string }
	// JSON -> Go struct
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n\n", titles) // [{Casablanca} {Cool Hand Luke} {Builitt}]

	var years []struct {
		Year int `json:"released"`
	}
	// JSON -> Go struct
	if err := json.Unmarshal(data, &years); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n\n", years) // [{1942} {1967} {1968}]

	var actors []struct{ Actors []string }
	// JSON -> Go struct
	if err := json.Unmarshal(data, &actors); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(actors) // [{"Humphrey Bogart", "Ingrid Bergman"} {"Paul Newman"} {"Steve McQueen", "Jacqueline Bisset"}]
}
