package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

// Fetch the json from the weapons/skins api
func getJson() []byte {
	// Read the cache
	bytes, err := readCache()

	// If the file is too old, request new data from the API
	if err != nil {
		if err.Error() == "cached file unusable: older than 48h" || strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println(err.Error())
			fmt.Println("Cache file unusable. Requesting new file from https://valorant-api.com/v1/weapons")

			// Request new data from the API
			resp, err := http.Get("https://valorant-api.com/v1/weapons")
			if err != nil {
				log.Fatalln(err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			bytes = body
			// Cache and return the new data
			writeCache(bytes)
			return bytes
		} else {
			log.Fatalln(err)
		}
	}

	return bytes
}

func main() {
	router := gin.Default()
	router.GET("/api/v1/phantom", ginmain)
	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func ginmain(c *gin.Context) {
	var response jsonData
	bytes := getJson()

	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
		return
	}

	data, err := generateJson(4, response)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, data)
}

func generateJson(i int, response jsonData) (skinResp, error) {

	if i > 9 {
		return skinResp{}, errors.New("index must be =< 9")
	}

	var resp []map[string]string
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

		resp = append(resp, map[string]string{response.Data[i].Skins[o].Name: icon})

	}

	return skinResp{resp}, nil
}
