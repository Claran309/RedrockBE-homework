package routes

import (
	"Redrock-lesson1/Redrock-lesson3/lvX/GinProject_LoggingAndRegisting/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//根目录：JWT中间件判断是否登录

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	//受保护路径：JWT中间件判断访问权限

	err := r.Run()
	if err != nil {
		panic("Failed to start Gin server: " + err.Error())
	}
}

//http://localhost:8080/register
