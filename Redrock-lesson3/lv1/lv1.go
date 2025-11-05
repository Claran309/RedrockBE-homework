package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/talk", func(c *gin.Context) {
		msg := c.Query("msg")
		var data string
		if msg == "ping" {
			data = "pong"
		} else if msg == "helloserver" {
			data = "helloclient"
		} else {
			data = "no information"
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	err := r.Run("")
	if err != nil {
		panic(err)
	}
}
