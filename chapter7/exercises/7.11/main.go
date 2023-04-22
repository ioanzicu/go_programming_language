// Add additional handlers so that clients can create, read, update, and delete
// database entries. For example, a request of the form /update?item=socks&price=6 will
// update the price of an item in the inventory and report an error if the item does not exist or if
// the price is invalid. (Warning: this change introduces concurrent variable updates.)

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		msg := fmt.Sprintf("item cannot be empty: %s\n", req.URL)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	price := req.URL.Query().Get("price")
	priceFloat32, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("cannot parse price: %s\n", err)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("item %s already exists in db\n", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = dollars(priceFloat32)
	w.WriteHeader(http.StatusCreated)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		msg := fmt.Sprintf("no such item: %s\n", item)
		http.Error(w, msg, http.StatusNotFound) // 404
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		msg := fmt.Sprintf("item cannot be empty: %s\n", req.URL)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	price := req.URL.Query().Get("price")
	priceFloat32, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("cannot parse price: %s\n", err)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("cannot update not existing item: %s\n", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	db[item] = dollars(priceFloat32)
	w.WriteHeader(http.StatusOK)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		msg := fmt.Sprintf("item cannot be empty: %s\n", req.URL)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("cannot delete not existing item: %s\n", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}

	delete(db, item)
	w.WriteHeader(http.StatusOK)
}

func main() {
	db := database{
		"shoes":   50,
		"socks":   5,
		"bananas": 33,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	fmt.Println("Start...")

	// net/http provides a global ServeMux instance called DefaultServeMux
	// and package-level functions called http.Handle and http.HandleFunc. To use Default-
	// ServeMux as the ser ver’s main handler, we needn’t pass it to ListenAndServe; nil will do.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
