package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func beefSummaryHandler(c *gin.Context) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch data"})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read data"})
		return
	}

	beefCount := countBeef(string(body))

	responseData := map[string]interface{}{
		"beef": beefCount,
	}

	c.JSON(http.StatusOK, responseData)
}

func main() {
	router := gin.Default()
	router.GET("/beef/summary", beefSummaryHandler)
	router.Run(":8080")
}
