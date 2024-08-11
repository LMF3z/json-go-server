package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func ExistJsonFile(path string) (bool, error) {

	file, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, err
	}

	if err != nil {
		return false, err
	}

	isJson := strings.Split(file.Name(), ".")[1]

	if isJson != "json" {
		log.Fatalln("File is not a JSON")
	}

	return true, nil

}

func ReadJsonFile(path string) (map[string]interface{}, error) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println("error to read file")
		return nil, err
	}

	defer file.Close()

	var data map[string]interface{}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println("Error al decodificar")
		return nil, err
	}

	return data, err
}
