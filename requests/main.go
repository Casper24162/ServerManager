package requests

import (
	"fmt"
	"log"
	"net/http"

	f "github.com/Casper24162/ServerManager/format"
	"github.com/Casper24162/ServerManager/jsonf"
)

type jsonDataStruct struct {
	ServerType    string `json:"Server-Type"`
	ServerVersion string `json:"Server-Version"`
}

func MakeServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println(f.Format("yellow", fmt.Sprintf("Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusUnsupportedMediaType)))
		writeResponse(w, http.StatusUnsupportedMediaType)
		return
	}

	var jsonData jsonDataStruct
	errDecodeJSON := jsonf.DecodeJSON(r.Body, &jsonData)
	if errDecodeJSON != nil {
		log.Println(f.Format("yellow", fmt.Sprintf("Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, http.StatusBadRequest)
		return
	}

	switch {
	case jsonData.ServerType == "":
		log.Println(f.Format("yellow", fmt.Sprintf("Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, http.StatusBadRequest)
		return
	case jsonData.ServerVersion == "":
		log.Println(f.Format("yellow", fmt.Sprintf("Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, http.StatusBadRequest)
		return
	}

	log.Println(f.Format("green", fmt.Sprintf("Got request from %v to create server of type %v:%v", r.RemoteAddr, jsonData.ServerType, jsonData.ServerVersion)))

	// Unzip tar/gz server file and execute 'exec.sh'

	writeResponse(w, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, code int) {
	// Add support for json messages
	w.WriteHeader(code)
	w.Write(nil)
}
