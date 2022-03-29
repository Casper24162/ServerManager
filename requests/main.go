package requests

import (
	"log"
	"net/http"
)

func MakeServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request from %v to create a server of type %v:%v", r.RemoteAddr, r.Header.Get("Server-Type"), r.Header.Get("Server-Version"))

	// Unzip tar/gz server file and execute 'exec.sh'
}
