package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// db := database{"shoes": 50, "socks": 5}
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	http.ListenAndServe("localhost:8000", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.HasPrefix(req.URL.Path, "/a") {
			muxA().ServeHTTP(w, req)
			return
		}

		if strings.HasPrefix(req.URL.Path, "/b") {
			muxB().ServeHTTP(w, req)
			return
		}
	}))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func muxA() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/a/list", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s %s\n", "/a", "/list")
	}))
	return mux
}

func muxB() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/b/list", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s %s\n", "/b", "/list")
	}))
	return mux
}
