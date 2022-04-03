package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	f "github.com/Casper24162/ServerManager/format"
	"github.com/Casper24162/ServerManager/jsonf"
)

type requestMake struct {
	ServerType    string `json:"server-type"`
	ServerVersion string `json:"server-version"`
}

type response struct {
	Message string `json:"message"`
}

func MakeServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusUnsupportedMediaType)))
		writeResponse(w, "'Content-Type' isn't 'application/json'", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData requestMake
	errDecodeJSON := jsonf.DecodeJSON(r.Body, &jsonData)
	if errDecodeJSON != nil {
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, "Body can't be decoded, please check if your body is formatted correctly.", http.StatusBadRequest)
		return
	}

	switch {
	case jsonData.ServerType == "":
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, "'server-type' can't be nil.", http.StatusBadRequest)
		return
	case jsonData.ServerVersion == "":
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, "'server-version' can't be nil.", http.StatusBadRequest)
		return
	}

	log.Println(f.Format("green", fmt.Sprintf("● Got request from %v to create server of type %v:%v", r.RemoteAddr, jsonData.ServerType, jsonData.ServerVersion)))

	// Unzip tar/gz server file and execute 'exec.sh'

	writeResponse(w, "Succes!", http.StatusOK)
}

func writeResponse(w http.ResponseWriter, message string, code int) {
	jsonStruct := response{Message: message}
	encodeErr := json.NewEncoder(w).Encode(jsonStruct)
	if encodeErr != nil {
		log.Println(f.Format("red", "● Internal server error when encoding response!"))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
	}

	w.WriteHeader(code)
	w.Write(nil)
}
