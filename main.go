package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"json-go-server/helpers"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var (
		path string
		port string
	)

	flag.StringVar(&path, "path", "pathFile is required", "pathFile is required")
	flag.StringVar(&port, "port", "8000", "default port is 8000")

	flag.Parse()

	_, err := helpers.ExistJsonFile(path)

	if err != nil {
		log.Fatal()
	}

	jsonFileData, err := helpers.ReadJsonFile(path)

	if err != nil {
		log.Fatal()
	}

	for key, value := range jsonFileData {

		endpointJsonData, err := helpers.ConvertArrJsonToInterface(value)

		if err != nil {
			log.Fatal("Internal server error")
			return
		}

		http.HandleFunc("/"+key, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			switch r.Method {
			case http.MethodGet:
				{
					jsonBytes, err := json.Marshal(endpointJsonData)

					if err != nil {
						http.Error(w, "Endpoint no valid", http.StatusBadRequest)
						return
					}

					w.Write(jsonBytes)
				}

			case http.MethodPost:
				{
					body, err := helpers.ConvertIoReadCloser(r.Body)

					if err != nil {
						res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
							"message": "Body is required",
						})

						w.Write(res)
					}

					body["id"] = strconv.Itoa(len(endpointJsonData) + 1)

					endpointJsonData = append(endpointJsonData, body)

					res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
						"message": "Item add successfully",
					})

					w.Write(res)

				}

			default:
				{

					res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
						"message": "nothing to show",
					})

					w.Write(res)
				}
			}

		})

		http.HandleFunc("/"+key+"/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			idParam := r.URL.Path
			parameters := strings.Split(idParam, "/")

			if len(parameters) > 3 {
				http.Error(w, "Not valid route", http.StatusBadRequest)
				return
			}
			idForSearch := parameters[2]

			switch r.Method {
			case http.MethodGet:
				{

					var itemFound []byte

					for _, value2 := range endpointJsonData {
						if value2["id"] == idForSearch {
							jsonDataRes, _ := helpers.ConvertStrToJson(value2)
							itemFound = jsonDataRes
						}
					}

					if len(itemFound) == 0 {
						http.Error(w, "Item by id "+idForSearch+" not found", http.StatusNotFound)
						return
					}

					w.Write(itemFound)

				}

			case http.MethodPatch:
				{

					body, err := helpers.ConvertIoReadCloser(r.Body)

					if err != nil {
						res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
							"message": "Body is required",
						})

						w.Write(res)
					}

					for index, val := range endpointJsonData {

						if val["id"] == idForSearch {
							for key, val2 := range body {
								endpointJsonData[index][key] = val2
								break
							}
						}

					}

					res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
						"message": "Item updated successfully",
					})

					w.Write(res)

				}

			case http.MethodDelete:
				{
					var newItems []map[string]interface{}

					for _, val := range endpointJsonData {
						if val["id"] != idForSearch {
							newItems = append(newItems, val)
						}
					}

					endpointJsonData = newItems

					res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
						"message": "Item deleted successfully",
					})

					w.Write(res)
				}

			default:
				{

					res, _ := helpers.ConvertMsgToBytes(map[string]interface{}{
						"message": "nothing to show",
					})

					w.Write(res)
				}
			}

		})
	}

	fmt.Println("dev api listen on port: ", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
