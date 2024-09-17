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

func main() {
	// Set the port for the server to run on
	var port uint16 = 8080

	// Comment out the below to enable debug mode
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Rifles
	router.GET("/api/v1/skin/vandal", handleVandal)
	router.GET("/api/v1/skin/bulldog", handleBulldog)
	router.GET("/api/v1/skin/phantom", handlePhantom)
	router.GET("/api/v1/skin/guardian", handleGuardian)

	// LMGs
	router.GET("/api/v1/skin/odin", handleOdin)
	router.GET("/api/v1/skin/ares", handleAres)

	// SMGs
	router.GET("/api/v1/skin/spectre", handleSpectre)
	router.GET("/api/v1/skin/stinger", handleStinger)

	// Shotguns
	router.GET("/api/v1/skin/judge", handleJudge)
	router.GET("/api/v1/skin/bucky", handleBucky)

	// Sidearms
	router.GET("/api/v1/skin/frenzy", handleFrenzy)
	router.GET("/api/v1/skin/ghost", handleGhost)
	router.GET("/api/v1/skin/sheriff", handleSheriff)
	router.GET("/api/v1/skin/classic", handleClassic)
	router.GET("/api/v1/skin/shorty", handleShorty)

	// Snipers
	router.GET("/api/v1/skin/operator", handleOP)
	router.GET("/api/v1/skin/outlaw", handleOutlaw)
	router.GET("/api/v1/skin/marshal", handleMarshal)

	// Knife
	router.GET("/api/v1/skin/melee", handleMelee)

	// Endpoint for all chromas
	router.GET("api/v1/chromas", handleChromas)

	// Search endpoint
	router.GET("/api/v1/search", search)

	fmt.Printf("Starting server at 0.0.0.0:%d\n", port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalln(err)
	}
}

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

func ginMain(index int) []skinResp {
	var response jsonData
	bytes := getJson()

	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
		return []skinResp{}
	}

	data, err := generateJson(index, response)
	if err != nil {
		panic(err)
	}

	return data
}

func generateJson(i int, response jsonData) ([]skinResp, error) {

	if i > 18 {
		return []skinResp{}, errors.New("index must be =< 18")
	}

	// Sort the slice alphabetically
	sort.Slice(response.Data[i].Skins, func(ii, j int) bool {
		return response.Data[i].Skins[ii].Name < response.Data[i].Skins[j].Name
	})

	var resp []skinResp
	for o := 0; o != len(response.Data[i].Skins); o++ {
		// Make sure the "Random Favorite Skin" does not appear in the response
		if !strings.Contains(response.Data[i].Skins[o].Name, "Random") {
			var icon string
			// Ensure a blank URL is never returned
			if response.Data[i].Skins[o].Icon == "" {
				icon = response.Data[i].Skins[o].Chromas[0].Icon
			} else {
				icon = response.Data[i].Skins[o].Icon
			}

			resp = append(resp, skinResp{
				Name: response.Data[i].Skins[o].Name,
				UUID: response.Data[i].Skins[o].UUID,
				Icon: icon})

		}
	}

	return resp, nil
}
