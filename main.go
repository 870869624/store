package main

import (
	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/api/authentication"
	"jinghaijun.com/store/api/catalogue"
	"jinghaijun.com/store/api/product"
	"jinghaijun.com/store/api/user"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/Authorization", authentication.SignIn) //需要获取token登陆所以不适合写进user，用户登陆（应该单独写个验证链接登陆）
	r.POST("/user", user.SignUp)                    //用户注册	r.POST("/user", user.SignUp)

	//将用户相关操作放进这个group
	Group_User := r.Group("/user")
	{
		//Group_User.Use(authorization.Auth)
		Group_User.GET("/:id", user.ListOne)         //用户查询全部信息
		Group_User.DELETE("/:id", user.Cancelletion) //用户注销
		Group_User.PATCH("", user.Update)
	}
	//需要获取验证码进行操作的步骤
	// Group_Code := r.Group("/code")
	// {
	// 	Group_Code.Use("", Ver)
	// }

	//关于商品的操作
	Group_Product := r.Group("/product")
	{
		// Group_Product.Use(authorization.Auth)
		Group_Product.POST("", product.Create)
		Group_Product.DELETE("/:id", product.DELETE)
		Group_Product.PATCH("", product.Update)
		Group_Product.GET("", product.List)
	}

	//关于分类的操作
	Group_Catalogue := r.Group("/catalogue")
	{
		//Group_Catalogue.Use(authentication.SignIn)
		Group_Catalogue.POST("", catalogue.Create)
		Group_Catalogue.GET("", catalogue.Listall)
		Group_Catalogue.GET("/:id", catalogue.ListOne)
		Group_Catalogue.PATCH("")
	}
	r.Run(":3000")

}
