package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.StaticFile("/cat.jpg", "./Redrock-lesson3/lv2/static/cat.jpg")

	err := r.Run("")
	if err != nil {
		panic(err)
	}
}
