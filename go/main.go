package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("sampledata.json")
	check(err)
	fmt.Print(string(dat))

	var unmarshal []map[string]interface{}
	json.Unmarshal(dat, &unmarshal)
	fmt.Print(unmarshal)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",

			"https://rankedrace.win",
			"https://www.rankedrace.win",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true, // only if using cookies
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/rank", func(c *gin.Context) {
		c.JSON(200, unmarshal)
	})

	router.Run() // listens on 0.0.0.0:8080 by default
}
