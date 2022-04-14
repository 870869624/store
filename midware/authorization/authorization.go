package authorization

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/token"
)

//获取Header中的token信息
func Auth(context *gin.Context) {
	tokenString := context.Request.Header.Get("Authorization") //获取授权书,从Header获取
	if tokenString == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "请登录",
		})
		return
	}
	claim, e := token.Parse(tokenString)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "请登录！",
		})
		return
	}
	context.Request.Header.Set("x-consumer-id", fmt.Sprintf("%d", claim.Id))
	context.Request.Header.Set("x-consumer-id", fmt.Sprintf("%d", claim.Kind))
}
