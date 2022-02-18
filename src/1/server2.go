package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	countResponse := make(map[string]int)
	mu.Lock()
	countResponse["count"] = count
	data, err := json.Marshal(countResponse)
	mu.Unlock()
	if err != nil {
		fmt.Fprintf(w, "{\"message\":\"Failed to read counter\"}")
	} else {
		fmt.Fprintf(w, "%s", string(data))
	}
}