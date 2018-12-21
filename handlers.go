package main

import (
    "fmt"
    "net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Showing users list")
}