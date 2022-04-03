package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Casper24162/ServerManager/config"
	f "github.com/Casper24162/ServerManager/format"
	"github.com/Casper24162/ServerManager/requests"
)

func main() {
	// Working OK, really proud of the for loop making the FH XD
	errMakeFiles := config.MakeFH()
	if errMakeFiles != nil {
		log.Fatalln(errMakeFiles)
	}

	var configData config.Config
	configData, errReadConfig := config.ReadConfig("data/config.json")
	if errReadConfig != nil {
		log.Fatalln(errReadConfig)
	}

	// HTTP Endpoints
	http.HandleFunc("/make", requests.MakeServer)

	log.Println(f.Format("green", fmt.Sprintf("● Server running on %v:%v", configData.Address, fmt.Sprint(configData.Port))))
	http.ListenAndServe(fmt.Sprintf("%v:%v", configData.Address, fmt.Sprint(configData.Port)), nil)
}
