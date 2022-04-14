package product

import (
	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/product"
)

func CreateProduct(context *gin.Context) {
	var product product.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"message": "参数错误！",
		})
		return
	}
	error := product.Create()

}
