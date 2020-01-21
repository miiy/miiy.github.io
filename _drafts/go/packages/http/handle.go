package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func helloHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main()  {

	http.Handle("/count", new(countHandler))

	http.HandleFunc("/hello", helloHandler)

	helloHandler2 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello2")
	}
	http.HandleFunc("/hello2", helloHandler2)

	http.HandleFunc("/hello3", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello3")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
/count
/hello
/hello2
/hello3
*/
