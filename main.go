package main

import (
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
	router.GET("/api/v1/skin/vandal", func(c *gin.Context) { handleSkins(c, 2) })
	router.GET("/api/v1/skin/bulldog", func(c *gin.Context) { handleSkins(c, 3) })
	router.GET("/api/v1/skin/phantom", func(c *gin.Context) { handleSkins(c, 4) })
	router.GET("/api/v1/skin/guardian", func(c *gin.Context) { handleSkins(c, 13) })

	// LMGs
	router.GET("/api/v1/skin/odin", func(c *gin.Context) { handleSkins(c, 0) })
	router.GET("/api/v1/skin/ares", func(c *gin.Context) { handleSkins(c, 1) })

	// SMGs
	router.GET("/api/v1/skin/spectre", func(c *gin.Context) { handleSkins(c, 16) })
	router.GET("/api/v1/skin/stinger", func(c *gin.Context) { handleSkins(c, 17) })

	// Shotguns
	router.GET("/api/v1/skin/judge", func(c *gin.Context) { handleSkins(c, 5) })
	router.GET("/api/v1/skin/bucky", func(c *gin.Context) { handleSkins(c, 6) })

	// Sidearms
	router.GET("/api/v1/skin/classic", func(c *gin.Context) { handleSkins(c, 8) })
	router.GET("/api/v1/skin/shorty", func(c *gin.Context) { handleSkins(c, 11) })
	router.GET("/api/v1/skin/frenzy", func(c *gin.Context) { handleSkins(c, 7) })
	router.GET("/api/v1/skin/ghost", func(c *gin.Context) { handleSkins(c, 9) })
	router.GET("/api/v1/skin/sheriff", func(c *gin.Context) { handleSkins(c, 10) })

	// Snipers
	router.GET("/api/v1/skin/operator", func(c *gin.Context) { handleSkins(c, 12) })
	router.GET("/api/v1/skin/outlaw", func(c *gin.Context) { handleSkins(c, 14) })
	router.GET("/api/v1/skin/marshal", func(c *gin.Context) { handleSkins(c, 15) })

	// Knife
	router.GET("/api/v1/skin/melee", func(c *gin.Context) { handleSkins(c, 18) })

	// Cards, Chromas, and Search endpoints
	router.GET("/api/v1/cards", handlePlayerCards)
	router.GET("api/v1/chromas", handleChromas)
	router.GET("/api/v1/search", search)

	fmt.Printf("Starting server at 0.0.0.0:%d\n", port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalln(err)
	}
}

func getJson(t string) []byte {
	// Read the cache
	bytes, err := readCache(t)
	// If the file is too old, request new data from the API
	if err != nil {
		if err.Error() == "cached file unusable: older than 48h" || strings.Contains(err.Error(), "no such file or directory") {
			// Use a different endpoint depending on
			// what we want to request
			var endpoint string
			switch t {
			case "skin":
				endpoint = "weapons/"
			case "card":
				endpoint = "playercards/"
			case "buddy":
				endpoint = "buddies/"
			default:
				log.Fatalf("%v is not a valid item to request.", t)
			}

			fmt.Println("Cached file unusable or not found. Requesting new file from https://valorant-api.com/v1/" + endpoint)

			// Request new data from the API
			resp, err := http.Get("https://valorant-api.com/v1/" + endpoint)
			if err != nil {
				log.Fatalln(err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			bytes = body
			// Cache and return the new data
			writeCache(t, bytes)
			return bytes
		} else {
			log.Fatalln(err)
		}
	}
	return bytes
}

func generateJson(i int, response skinJsonData) []skinResp {

	// Sort the slice alphabetically
	sort.Slice(response.Data[i].Skins, func(ii, j int) bool {
		return response.Data[i].Skins[ii].Name < response.Data[i].Skins[j].Name
	})

	var resp []skinResp
	// Iterate over all the skins
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

			// Append each skin to the response
			resp = append(resp, skinResp{
				Name: response.Data[i].Skins[o].Name,
				UUID: response.Data[i].Skins[o].UUID,
				Icon: icon})

		}
	}

	return resp
}
