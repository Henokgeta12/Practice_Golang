package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    // Create an HTTP client
    client := &http.Client{}

    // Define the URL to request
    url := "https://jsonplaceholder.typicode.com/posts/1"

    // Send an HTTP GET request
    resp, err := client.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close() // Ensure the response body is closed after the function completes


    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the response
    fmt.Println("Response:", string(body))
}