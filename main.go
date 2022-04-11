package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/user", SignUp)          //用户注册
	r.POST("/Authorization", SignIn) //用户登陆
}
