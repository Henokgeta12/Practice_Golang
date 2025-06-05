package main

import (
    "fmt"
    "net/http"
)

// Define a struct
type HomeHandler struct{}

// Implement the ServeHTTP method for the HomeHandler struct
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the home page!")
}

func main() {
    // Create a new ServeMux
    mux := http.NewServeMux()

    // Register an instance of HomeHandler as the handler for the root path "/"
    homeHandler := HomeHandler{}
    mux.Handle("/", homeHandler)

    // Start the HTTP server
    http.ListenAndServe(":8080", mux)
}
