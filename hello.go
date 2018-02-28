package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func vipin(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Created by vipin")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/vipin", vipin)
    http.ListenAndServe(":80", nil)
}