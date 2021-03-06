package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!\nYour http request method is %s\n", r.Method)
}

func main() {
    http.HandleFunc("/", helloHandler)
    fmt.Println("Serving")
    http.ListenAndServe(":8080", nil)
}
