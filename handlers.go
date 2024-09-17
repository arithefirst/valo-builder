package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func handleChromas(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(http.StatusNotFound, errResp{Error: "A UUID Must be specified.", Status: http.StatusNotFound})
		return
	}

	var response jsonData
	bytes := getJson()

	err := json.Unmarshal(bytes, &response)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i != len(response.Data); i++ {
		for _, v := range response.Data[i].Skins {
			if v.UUID == uuid {
				c.JSON(http.StatusOK, v)
				return
			}
		}
	}

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
