package config

import (
	"encoding/json"
	"os"

	"github.com/Casper24162/ServerManager/jsonf"
)

type Config struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func MakeFH() error {
	dirArray := [...]string{"data", "data/servers", "data/servers/active", "data/servers/archives"}

	for i := 0; i < len(dirArray); i++ {
		errMakeDir := os.Mkdir(dirArray[i], 0660)

		if errMakeDir != nil && os.IsNotExist(errMakeDir) {
			return errMakeDir
		}
	}

	if _, errCheckExists := os.Stat("data/config.json"); errCheckExists == nil {
		return nil
	}

	// Default config data
	config := Config{Port: 2424, Address: "localhost"}

	configBytes, errMarshal := json.Marshal(config)
	if errMarshal != nil {
		return errMarshal
	}

	if errWriteConfig := os.WriteFile("data/config.json", configBytes, 0660); errWriteConfig != nil {
		return errWriteConfig
	}

	return nil
}

func ReadConfig(path string) (Config, error) {
	file, errOpenFile := os.Open(path)
	defer file.Close()
	if errOpenFile != nil {
		return Config{}, errOpenFile
	}

	var configStruct Config
	errDecodeJSON := jsonf.DecodeJSON(file, &configStruct)
	if errDecodeJSON != nil {
		return Config{}, errDecodeJSON
	}

	return configStruct, nil
}
