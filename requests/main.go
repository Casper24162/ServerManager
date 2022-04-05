package requests

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	f "github.com/Casper24162/ServerManager/format"
	"github.com/Casper24162/ServerManager/jsonf"
)

type requestMake struct {
	ServerType    string `json:"server-type"`
	ServerVersion string `json:"server-version"`
}

type response struct {
	Message  string `json:"message"`
	ServerID string `json:"server"`
}

func MakeServer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusUnsupportedMediaType)))
		writeResponse(w, response{"'Content-Type' isn't 'application/json'", "NULL"}, http.StatusUnsupportedMediaType)
		return
	}

	var jsonData requestMake
	errDecodeJSON := jsonf.DecodeJSON(r.Body, &jsonData)
	if errDecodeJSON != nil {
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, response{"Body can't be decoded, please check if your body is formatted correctly.", "NULL"}, http.StatusBadRequest)
		return
	}

	switch {
	case jsonData.ServerType == "":
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, response{"'server-type' can't be nil.", "NULL"}, http.StatusBadRequest)
		return
	case jsonData.ServerVersion == "":
		log.Println(f.Format("yellow", fmt.Sprintf("● Got bad request from %v, returned HTTP:%v", r.RemoteAddr, http.StatusBadRequest)))
		writeResponse(w, response{"'server-version' can't be nil.", "NULL"}, http.StatusBadRequest)
		return
	}

	log.Println(f.Format("green", fmt.Sprintf("● Got request from %v to create server of type %v:%v", r.RemoteAddr, jsonData.ServerType, jsonData.ServerVersion)))

	// Unzip tar/gz server file and execute 'exec.sh'

	writeResponse(w, response{"Succes!", genServerID(6)}, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, jsonStruct response, code int) {
	w.WriteHeader(code)

	encodeErr := json.NewEncoder(w).Encode(jsonStruct)
	if encodeErr != nil {
		log.Println(f.Format("red", "● Internal server error when encoding response!"))
	}
}

func genServerID(length int) string {
	rand.Seed(time.Now().UnixMilli())
	chars := []string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var id string
	for i := 0; i <= length; i++ {
		id = id + chars[rand.Intn(len(chars))]
	}
	return id
}
