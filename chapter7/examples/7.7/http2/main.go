package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			// w.WriteHeader(http.StatusNotFound) // 404
			// fmt.Fprintf(w, "no such item: %q\n", item)
			// OR
			msg := fmt.Sprintf("no such page: %s\n", req.URL)
			http.Error(w, msg, http.StatusNotFound) // 404

			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database{
		"shoes":   50,
		"socks":   5,
		"bananas": 33,
	}
	fmt.Println("Start...")
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
