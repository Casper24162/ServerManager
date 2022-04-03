package jsonf

import (
	"encoding/json"
	"io"
)

func DecodeJSON(jsonData io.Reader, jsonStruct any) error {
	decoder := json.NewDecoder(jsonData)
	decoder.DisallowUnknownFields()

	decodeErr := decoder.Decode(&jsonStruct)
	if decodeErr != nil {
		return decodeErr
	}
	return nil
}
