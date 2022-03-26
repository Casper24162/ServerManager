package main

import (
	f "fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":2424", nil))
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	f.Println("Hello World!")
}
