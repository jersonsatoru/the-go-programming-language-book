package webserver

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func Webserver() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count += 1
	mu.Unlock()

	for key, value := range r.Header {
		fmt.Fprintf(w, "%s - %s", key, value)
	}

	fmt.Fprintf(w, "HOST - %s", r.Host)
	fmt.Fprintf(w, "REMOTE_ADDR - %s", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range r.Form {
		fmt.Fprintf(w, "%s - %s", key, value)
	}

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, fmt.Sprintf("%d", count))
	mu.Unlock()
}
