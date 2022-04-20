package catalogue

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/catalogue"
)

func Create(context *gin.Context) {
	var cata catalogue.Catalogue
	if err := context.ShouldBindJSON(&cata); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"提示": "参数错误",
		})
		return
	}
	e := cata.Creat()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": e.Error(),
		})
		return
	}
	context.AbortWithStatusJSON(http.StatusCreated, gin.H{
		"message": "创建成功",
	})
}

//商品分类的查看方法，也是从uri中获取参数然后通过联合查找额度方式查找出来
func Listall(context *gin.Context) {
	order_by := context.QueryArray("order")
	page := context.DefaultQuery("page", "1")
	number := context.DefaultQuery("number", "5")
	ipage, e := strconv.Atoi(page)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "page为数字类型",
		})
		return
	}
	inumber, e := strconv.Atoi(number)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "number为数字类型",
		})
		return
	}
	pagnition := catalogue.Pagination{
		Page:   ipage,
		Number: inumber,
	}
	order := catalogue.Order{
		By: order_by,
	}
	query := catalogue.CatalogueQuery{
		Pagination: pagnition,
		Order:      order,
	}
	response, e := catalogue.Listall(&query)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	context.AbortWithStatusJSON(http.StatusAccepted, response)
}

//展示类目中的全部商品
func ListOne(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	number := context.DefaultQuery("number", "5")
	order_by := context.QueryArray("order")
	title := context.Query("tilte")
	ipage, e := strconv.Atoi(page)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "page为数字类型",
		})
		return
	}
	inumber, e := strconv.Atoi(number)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "number为数字类型",
		})
		return
	}
	pagnition := catalogue.Pagination{
		Page:   ipage,
		Number: inumber,
	}
	order := catalogue.Order{
		By: order_by,
	}
	query := catalogue.CatalogueQuery{
		Pagination: pagnition,
		Order:      order,
		Title:      title,
	}
	response, e := catalogue.ListOne(&query)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	context.AbortWithStatusJSON(http.StatusOK, response)
}

//更新目录
func Update(context *gin.Context) {
	var catalogue catalogue.Catalogue
	if err := context.ShouldBindQuery(&catalogue); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
	}
	e := catalogue.Update()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e)
		return
	}
}
