package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/about", helpHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world! :)")
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Woah about page")
}