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

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
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
}

func main() {
	db := database{
		"shoes":   50,
		"socks":   5,
		"bananas": 33,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	fmt.Println("Start...")

	// net/http provides a global ServeMux instance called DefaultServeMux
	// and package-level functions called http.Handle and http.HandleFunc. To use Default-
	// ServeMux as the ser ver’s main handler, we needn’t pass it to ListenAndServe; nil will do.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
