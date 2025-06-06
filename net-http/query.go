package main

import (
	"fmt"
	"net/http"
)

//http://localhost:7070/query?name=abc&email=123@.com
func queryHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Fprintln(w, "path:", r.URL.Path)
	fmt.Fprintln(w, "raw query:", r.URL.RawQuery)
	fmt.Fprintln(w, "query:", q)
	fmt.Fprintln(w, "name:", q.Get("name"))
	fmt.Fprintln(w, "email:", q.Get("email"))
}

func main() {
	http.HandleFunc("/query", queryHandler)
	http.ListenAndServe(":7070", nil)
}