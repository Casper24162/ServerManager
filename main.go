package main

import (
	"log"
	"net/http"

	"github.com/Casper24162/ServerManager/requests"
)

func main() {
	go http.HandleFunc("/make", requests.MakeServer)

	log.Fatal(http.ListenAndServe(":2424", nil))
}
