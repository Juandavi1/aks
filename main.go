package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"container_id": id(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func id() string {
	id := os.Getenv("HOSTNAME")
	if len(id) > 0 {
		return id
	}
	return "there isnt id."
}
