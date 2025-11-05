package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentRequest struct {
	Name  string    `json:"name"`
	Score []float64 `json:"score"`
}

func AverageCalculate(c *gin.Context) {
	var student StudentRequest
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Information format is invalid",
		})
		return
	}

	sum := 0.0
	for _, value := range student.Score {
		sum += value
	}
	sum /= float64(len(student.Score))

	c.JSON(http.StatusOK, gin.H{
		"average": sum,
	})
}

func main() {
	r := gin.Default()
	
	r.POST("/scores", AverageCalculate)

	err := r.Run("")
	if err != nil {
		panic(err)
	}
}
