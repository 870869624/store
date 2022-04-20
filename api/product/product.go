package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/product"
)

func Create(context *gin.Context) {
	var product product.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"message": "参数错误！",
		})
		return
	}
	error := product.Create()
	context.AbortWithStatusJSON(400, error.Error())
}
func DELETE(context *gin.Context) {
	var product product.Product
	i := context.Param("id") //获取uri中的ID
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}
	v, e := strconv.Atoi(i)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	error := product.Delete(v)
	if error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, error.Error())
	}
	context.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"message": "删除成功",
	})
}
func List(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	number := context.DefaultQuery("number", "10")
	price_gt := context.Query("price_gt")
	price_lt := context.Query("price_lt")
	title := context.Query("title")
	catalogue_id := context.Query("catalogue_id")
	order_by := context.QueryArray("order")
	ipage, e := strconv.Atoi(page)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "page应该为数字类型",
		})
		return
	}
	inumber, e := strconv.Atoi(number)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "number应为数字类型",
		})
		return
	}
	pagination := product.Pagination{
		Page:   ipage,
		Number: inumber,
	}
	order := product.Order{
		By: order_by,
	}
	query := product.ProductsearchQuery{
		Pagination: pagination,
		Order:      order,
	}
	if price_gt != "" {
		i_price_gt, e := strconv.Atoi(price_gt)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "最小价格应为数字型",
			})
			return
		}
		query.Price_gt = i_price_gt
	}
	if price_lt != "" {
		i_price_lt, e := strconv.Atoi(price_lt)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "最小价格应为数字型",
			})
			return
		}
		query.Price_lt = i_price_lt
	}
	if title != "" {
		query.Title = title
	}
	if catalogue_id != "" {
		i_catalogue_id, e := strconv.Atoi(catalogue_id)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "类目id为数字类型",
			})
			return
		}
		query.Catalogue_ID = uint64(i_catalogue_id)
	}
	response, e := product.List(&query)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
		return
	}
	context.AbortWithStatusJSON(http.StatusAccepted, response)
}
func Update(context *gin.Context) {
	var product product.Product
	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误！",
		})
		return
	}
	e := product.Update()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	context.AbortWithStatusJSON(http.StatusAccepted, gin.H{
		"message": "修改成功",
	})
}
