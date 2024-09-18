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

	// If weaponInt > 18, return an error
	if weaponInt > 18 {
		c.JSON(http.StatusNotFound, errResp{Error: "Weapon must be >= 0 && <= 18", Status: http.StatusNotFound})
		return
	}

	// Get data from the cache/api
	var response jsonData
	bytes := getJson()

	// Unmarshal the json into the struct
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
	}

	// Sort the slice alphabetically
	sort.Slice(response.Data[weaponInt].Skins, func(ii, j int) bool {
		return response.Data[weaponInt].Skins[ii].Name < response.Data[weaponInt].Skins[j].Name
	})

	// Create skins arr
	var skins []skin

	// For every skin in specified weapon
	for i := 0; i != len(response.Data[weaponInt].Skins); i++ {
		// If lowercase weapon name contains lowercase query, add to skins arr
		if strings.Contains(strings.ToLower(response.Data[weaponInt].Skins[i].Name), strings.ToLower(query)) {
			skins = append(skins, response.Data[weaponInt].Skins[i])
		}
	}

	// If skins is not empty, return it
	if len(skins) != 0 {
		c.JSON(http.StatusOK, skins)
		return
	}

	// If skins empty, respone with no skins found
	c.JSON(http.StatusNotFound, errResp{Error: "The requested query was not found.", Status: http.StatusNotFound})
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
	var response jsonData
	bytes := getJson()

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

// Yes, I know this is really redundant
// It's just (IMO) the cleanest way to do it w/o affecting performance
func handleOdin(c *gin.Context) {
	data := ginMain(0)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleAres(c *gin.Context) {
	data := ginMain(1)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleVandal(c *gin.Context) {
	data := ginMain(2)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleBulldog(c *gin.Context) {
	data := ginMain(3)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handlePhantom(c *gin.Context) {
	data := ginMain(4)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleJudge(c *gin.Context) {
	data := ginMain(5)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleBucky(c *gin.Context) {
	data := ginMain(6)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleFrenzy(c *gin.Context) {
	data := ginMain(7)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleClassic(c *gin.Context) {
	data := ginMain(8)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleGhost(c *gin.Context) {
	data := ginMain(9)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleSheriff(c *gin.Context) {
	data := ginMain(10)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleShorty(c *gin.Context) {
	data := ginMain(11)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleOP(c *gin.Context) {
	data := ginMain(12)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleGuardian(c *gin.Context) {
	data := ginMain(13)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleOutlaw(c *gin.Context) {
	data := ginMain(14)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleMarshal(c *gin.Context) {
	data := ginMain(15)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleSpectre(c *gin.Context) {
	data := ginMain(16)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleStinger(c *gin.Context) {
	data := ginMain(17)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}

func handleMelee(c *gin.Context) {
	data := ginMain(18)
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, data)
}
