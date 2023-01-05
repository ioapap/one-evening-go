package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls []string
var stats = map[string]int{}

func main() {
	http.HandleFunc("/hello", myFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myFunction(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	calls = append(calls, name)
	stats[name]++

	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
	fmt.Fprint(w, "Hello, ", name)
}
