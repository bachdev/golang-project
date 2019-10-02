package helper

import (
	"os"

	jsoniter "github.com/json-iterator/go"
)

// Bind dữ liệu từ file json 
func DecodeDataFromJsonFile(f *os.File, data interface{}) error {
	jsonParser := jsoniter.NewDecoder(f)
	err := jsonParser.Decode(&data)
	if err != nil {
		return err
	}

	return nil
}
