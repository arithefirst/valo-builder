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
	router.GET("/api/v1/vandal", handleVandal)
	router.GET("/api/v1/bulldog", handleBulldog)
	router.GET("/api/v1/phantom", handlePhantom)
	router.GET("/api/v1/guardian", handleGuardian)

	// LMGs
	router.GET("/api/v1/odin", handleOdin)
	router.GET("/api/v1/ares", handleAres)

	// SMGs
	router.GET("/api/v1/spectre", handleSpectre)
	router.GET("/api/v1/stinger", handleStinger)

	// Shotguns
	router.GET("/api/v1/judge", handleJudge)
	router.GET("/api/v1/bucky", handleBucky)

	// Sidearms
	router.GET("/api/v1/frenzy", handleFrenzy)
	router.GET("/api/v1/ghost", handleGhost)
	router.GET("/api/v1/sheriff", handleSheriff)
	router.GET("/api/v1/classic", handleClassic)
	router.GET("/api/v1/shorty", handleShorty)

	// Snipers
	router.GET("/api/v1/operator", handleOP)
	router.GET("/api/v1/outlaw", handleOutlaw)
	router.GET("/api/v1/marshall", handleMarshall)

	// Knife
	router.GET("/api/v1/melee", handleMelee)

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

func ginMain(index int) skinResp {
	var response jsonData
	bytes := getJson()

	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
		return skinResp{}
	}

	data, err := generateJson(index, response)
	if err != nil {
		panic(err)
	}

	return data
}

func generateJson(i int, response jsonData) (skinResp, error) {

	if i > 18 {
		return skinResp{}, errors.New("index must be =< 18")
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
