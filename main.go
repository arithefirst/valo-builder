package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

// Fetch the json from the weapons/skins api
func getJson() []byte {
	resp, err := http.Get("https://valorant-api.com/v1/weapons")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	return body
}

func main() {
	var response jsonData
	bytes, err := readCache()

	// If the file is too old, request new data from the API
	if err != nil {
		if err.Error() == "cached file unusable: older than 48h" {
			fmt.Println("Cache file unusable. Requesting new file from https://valorant-api.com/v1/weapons")
			bytes = getJson()
			writeCache(bytes)
		} else {
			log.Fatalln(err)
		}
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
		return
	}

	var jsonString string
	data, err := generateJson(4, jsonString, response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)

}

func generateJson(i int, jsonOut string, response jsonData) (string, error) {

	if i > 9 {
		return "", errors.New("index must be =< 9")
	}

	jsonOut += "{\"Data\": ["
	for o := 0; o != len(response.Data[i].Skins); o++ {

		// Sort the slice alphabetically
		sort.Slice(response.Data[i].Skins, func(ii, j int) bool {
			return response.Data[i].Skins[ii].Name < response.Data[i].Skins[j].Name
		})

		var icon string
		// Ensure a blank URL is never returned
		if response.Data[i].Skins[o].Icon == "" {
			icon = response.Data[i].Skins[o].Chromas[0].Icon
		} else {
			icon = response.Data[i].Skins[o].Icon
		}

		jsonOut += fmt.Sprintf("{\"Name\":\"%v\",\"URL\":\"%v\"}", response.Data[i].Skins[o].Name, icon)
		if o != len(response.Data[i].Skins)-1 {
			jsonOut += ","
		}
	}

	jsonOut += "]}"

	return jsonOut, nil
}
