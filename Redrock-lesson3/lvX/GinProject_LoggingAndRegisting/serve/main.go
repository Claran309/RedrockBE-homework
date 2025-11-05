package main

//使用postman进行测试
import (
	"Redrock-lesson1/Redrock-lesson3/lvX/GinProject_LoggingAndRegisting/api/routes"
)

func main() {
	routes.InitRouter()
}

/*
注册时：
前端给:
username
password
email

登录时：
前端给：
login_key （用户名或邮箱）
password
*/
