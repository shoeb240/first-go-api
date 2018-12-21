package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func RequestHandle() {
    myRouter := mux.NewRouter()
    myRouter.HandleFunc("/users", getUsers).Methods("GET")

    log.Fatal(http.ListenAndServe(":8081", myRouter))
}