package main

import (
	"fmt"
	"net/http"
)
func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s %s\n", r.Method, r.URL.Path,r.Header)
		next.ServeHTTP(w, r)
	})
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Header mildeware")
}


func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.Handle("/", logMiddleware(http.HandlerFunc(hello)))
	http.HandleFunc("/header", headerMiddleware(http.HandlerFunc(header))
	http.ListenAndServe(":7070", nil)
}