package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func search(c *gin.Context) {
	// Get the weapon and query from the querystrings
	c.Header("Access-Control-Allow-Origin", "*")
	weapon := c.Query("weapon")
	query := c.Query("q")

	// If there is no query||weapon return 404
	if query == "" {
		c.JSON(http.StatusNotFound, errResp{Error: "Please enter a query.", Status: http.StatusNotFound})
		return
	} else if weapon == "" {
		c.JSON(http.StatusNotFound, errResp{Error: "Please set a weapon (int 0-18).", Status: http.StatusNotFound})
		return
	}

	// Convert weapon to int, if invalid return an error
	weaponInt, err := strconv.Atoi(weapon)
	if err != nil {
		c.JSON(http.StatusNotFound, errResp{Error: err.Error(), Status: http.StatusNotFound})
		return
	}

	// If weaponInt > 21 or < 0, return an error
	if weaponInt > 21 || weaponInt < 0 {
		c.JSON(http.StatusNotFound, errResp{Error: "Weapon must be >= 0 && <= 19", Status: http.StatusNotFound})
		return
	}

	// If weaponInt is 19, search for cards instead of skins
	if weaponInt == 19 {
		// Get data from the cache/api
		var response cardJsonData
		bytes := getJson("card")

		// Unmarshal the json into the struct
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatalln(err)
		}

		// Sort the slice alphabetically
		sort.Slice(response.Data, func(ii, j int) bool {
			return response.Data[ii].Name < response.Data[j].Name
		})

		// Create cards array
		var cards []card
		for i := 0; i != len(response.Data); i++ {
			// If lowercase weapon name contains lowercase query, add to skins arr
			if strings.Contains(strings.ToLower(response.Data[i].Name), strings.ToLower(query)) {
				cards = append(cards, response.Data[i])
			}
		}
		c.JSON(http.StatusOK, cards)
		return
	}

	// If weaponInt is 20, search for buddies instead of skins
	if weaponInt == 20 {
		// Get data from the cache/api
		var response buddyJsonData
		bytes := getJson("buddy")

		// Unmarshal the json into the struct
		err = json.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatalln(err)
		}

		// Sort the slice alphabetically
		sort.Slice(response.Data, func(ii, j int) bool {
			return response.Data[ii].Name < response.Data[j].Name
		})

		// Create buddies array
		var buddies []buddy
		for i := 0; i != len(response.Data); i++ {
			// If lowercase weapon name contains lowercase query, add to skins arr
			if strings.Contains(strings.ToLower(response.Data[i].Name), strings.ToLower(query)) {
				buddies = append(buddies, response.Data[i])
			}
		}
		c.JSON(http.StatusOK, buddies)
		return
	}

	// If weaponInt is 21, search for titles instead of skins
	if weaponInt == 21 {
		// Get the data from the cache/api
		var response titleJsonData
		bytes := getJson("title")

		// Unmarshal the json into the struct
		err := json.Unmarshal(bytes, &response)
		if err != nil {
			log.Fatalln(err.Error())
		}

		// Sort the slice alphabetically
		sort.Slice(response.Data, func(ii, j int) bool {
			return response.Data[ii].Content < response.Data[j].Content
		})

		// Create titles array
		var titles []title
		for i := 0; i != len(response.Data); i++ {
			// If lowercase weapon name contains lowercase query, add to skins arr
			if strings.Contains(strings.ToLower(response.Data[i].Content), strings.ToLower(query)) {
				titles = append(titles, response.Data[i])
			}
		}
		c.JSON(http.StatusOK, titles)
		return
	}

	// Get data from the cache/api
	var response skinJsonData
	bytes := getJson("skin")

	// Unmarshal the json into the struct
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
	}

	data := generateJson(weaponInt, response)

	// Sort the slice alphabetically
	sort.Slice(data, func(ii, j int) bool {
		return data[ii].Name < data[j].Name
	})

	// Create skins arr
	var skins []skinResp
	// For every skin in specified weapon
	for i := 0; i != len(data); i++ {
		// If lowercase weapon name contains lowercase query, add to skins arr
		if strings.Contains(strings.ToLower(data[i].Name), strings.ToLower(query)) {
			skins = append(skins, data[i])
		}
	}

	c.JSON(http.StatusOK, skins)
}

func handleChromas(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	uuid := c.Query("uuid")

	// If UUID empty return a 404
	if uuid == "" {
		c.JSON(http.StatusNotFound, errResp{Error: "A UUID Must be specified.", Status: http.StatusNotFound})
		return
	}

	// Get the data from the cache/api
	var response skinJsonData
	bytes := getJson("skin")

	// Unmarshal the json into the struct
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
	}

	// For each weapon
	for i := 0; i != len(response.Data); i++ {
		// For each skin in each weapon
		for _, v := range response.Data[i].Skins {
			if v.UUID == uuid {
				// If UUID Matches return the skin and chromas
				c.JSON(http.StatusOK, v)
				return
			}
		}
	}

	// If UUID not found, return a 404
	c.JSON(http.StatusNotFound, errResp{Error: "The requested UUID was not found.", Status: http.StatusNotFound})
}

func handleSkins(c *gin.Context, i int) {
	// Get the data from the cache/api
	var response skinJsonData
	bytes := getJson("skin")

	// Unmarshal the json into the struct
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Parse the data to sort alphabetically and to
	// ensure blank icons will never be given, and that
	// random favourite skin will never appear
	data := generateJson(i, response)

	// Return the data
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handlePlayerCards(c *gin.Context) {
	// Get the data from the cache/api
	var response cardJsonData
	bytes := getJson("card")

	// Unmarshal the json into the struct
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Sort the slice alphabetically
	sort.Slice(response.Data, func(ii, j int) bool {
		return response.Data[ii].Name < response.Data[j].Name
	})

	// Return the data
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response.Data)
}

func handleBuddies(c *gin.Context) {
	// Get the data from the cache/api
	var response buddyJsonData
	bytes := getJson("buddy")

	// Unmarshal the json into the struct
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Sort the slice alphabetically
	sort.Slice(response.Data, func(ii, j int) bool {
		return response.Data[ii].Name < response.Data[j].Name
	})

	// Return the data
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response.Data)
}

func handleTitles(c *gin.Context) {
	// Get the data from the cache/api
	var response titleJsonData
	bytes := getJson("title")

	// Unmarshal the json into the struct
	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Sort the slice alphabetically
	sort.Slice(response.Data, func(ii, j int) bool {
		return response.Data[ii].Content < response.Data[j].Content
	})

	// Return the data
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, response.Data)
}
