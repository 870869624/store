package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/authentication"
)

func SignIn(context *gin.Context) {
	var auth authentication.Authentication
	if err := context.ShouldBindJSON(&auth); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
	}
	token, err := auth.SignIn()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"token": token,
	})

}
