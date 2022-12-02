package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// mock db
var db = make(map[string]string)

func redirect(c *gin.Context) {
	shortUrl := c.Params.ByName("shortUrl")
	fmt.Println("shortenedUrl: ", shortUrl)

	longUrl := db[shortUrl]

	c.Redirect(301, string("https://"+longUrl))

}

func shortenUrl(c *gin.Context) {
	shortUrl := fmt.Sprint(rand.Int63n(1000))
	fmt.Println("shortenedUrl: ", shortUrl)

	longUrl := c.Params.ByName("longUrl")
	if longUrl != "" {
		db[shortUrl] = longUrl
	}

	c.Redirect(301, "/"+shortUrl)
}

func index(c *gin.Context) {
	if c.Params.ByName("shortUrl") != "" {
		c.JSON(200, gin.H{
			"shortenedUrl": "http://localhost:8080/redirect/" + c.Params.ByName("shortUrl"),
		})
	} else {
		c.JSON(200, gin.H{
			"welcome": "create a shortenedUrl",
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	r := gin.Default()

	r.GET("/:shortUrl", index)
	r.POST("/shortenUrl/:longUrl", shortenUrl)
	r.GET("/redirect/:shortUrl", redirect)

	r.Run()
}
