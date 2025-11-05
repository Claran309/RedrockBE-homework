package handlers

import (
	"Redrock-lesson1/Redrock-lesson3/lvX/GinProject_LoggingAndRegisting/dao"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Register(c *gin.Context) {
	//传入用户信息
	username, ok1 := c.GetPostForm("username")
	password, ok2 := c.GetPostForm("password")
	email, ok3 := c.GetPostForm("email")

	//并发安全
	dao.DataSync.Lock()
	defer dao.DataSync.Unlock()

	//信息不完整
	if !ok1 || !ok2 || !ok3 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "Information is incomplete",
		})
		return
	}

	//判断用户信息是否合法：
	//1. 用户名是否被使用过
	if flag := dao.CheckUsername(username); flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "user already exists",
		})
		return
	}

	//2. 密码时候否符合格式（仅包含英文字母和数字）
	var flagPassword bool
	for i := 0; i < len(password); i++ {
		if !((password[i] >= 'a' && password[i] <= 'z') || (password[i] >= '0' && password[i] <= '9') || (password[i] >= 'A' && password[i] <= 'Z')) {
			flagPassword = true
		}
	}
	if flagPassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "Password format is incorrect",
		})
		return
	}

	//3. 邮箱是否被注册过
	if flag := dao.CheckEmail(email); flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "Email already been registered",
		})
		return
	}

	//传出用户信息到数据库
	dao.AddUser(username, password, email)

	//JSON返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "User registered successfully",
	})
}

func Login(c *gin.Context) {
	loginKey := c.PostForm("login_key")
	password := c.PostForm("password")
	var username string

	//并发安全
	dao.DataSync.Lock()
	defer dao.DataSync.Unlock()

	//判断是邮箱登录还是用户名登录
	var at, point bool
	for i := 0; i < len(loginKey); i++ {
		if loginKey[i] == '@' {
			at = true
		}
		if loginKey[i] == '.' {
			point = true
		}
	}
	if at && point { // 邮箱登录
		username = dao.EmailToUsername(loginKey)
	} else { // 用户名登录
		username = loginKey
	}

	//检查用户是否存在
	if flag := dao.CheckUsername(username); !flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "User does not exist",
		})
		return
	}

	//检验密码正确性
	if password != dao.SelectPassword(username) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "Incorrect password",
		})
		return
	}

	//JSON返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "login successful",
	})
}
