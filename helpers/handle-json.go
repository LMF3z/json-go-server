package helpers

import (
	"encoding/json"
	"fmt"
	"io"
)

func ConvertArrJsonToInterface(value interface{}) ([]map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(value)

	if err != nil {
		fmt.Println("Error al leer string")
		return nil, err
	}

	var interfaceData []map[string]interface{}

	err = json.Unmarshal([]byte(jsonBytes), &interfaceData)

	if err != nil {
		fmt.Println("Error al convertir a interface string", err)
		return nil, err
	}

	return interfaceData, nil
}

func ConvertMsgToBytes(msg map[string]interface{}) ([]byte, error) {

	jsonBytes, err := json.Marshal(msg)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertStrToJson(value map[string]interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(value)

	if err != nil {
		fmt.Println("Error al leer string")
		return nil, err
	}

	return jsonBytes, nil
}

func ConvertIoReadCloser(b io.ReadCloser) (map[string]interface{}, error) {
	var data map[string]interface{}

	err := json.NewDecoder(b).Decode(&data)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, err
}
