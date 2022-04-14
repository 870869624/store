package verificationcode

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCode(context *gin.Context) {
	codeString := context.Request.Header.Get("VerificationCode")
	if codeString == "" {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"提示": "请验证手机号",
		})
	}
}
