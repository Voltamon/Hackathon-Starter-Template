package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// start embedded PocketBase
	StartPocketBase()

	// Allow all CORS for development
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/pocketbase", func(c *gin.Context) {
		// example: access PocketBase instance
		pb := GetPocketBase()
		pbStatus := "not running"
		if pb != nil {
			pbStatus = "running"
		}

		c.JSON(http.StatusOK, gin.H{"pocketbase": pbStatus})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	_ = r.Run(":8080")
}
