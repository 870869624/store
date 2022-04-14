package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"jinghaijun.com/store/models/user"
)

//用户登录 1.取得数据，然后然后插入数据库，检测数据库中是否已经存在用户的数据（中间需要进行加密）
func SignUp(context *gin.Context) {
	var user user.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	e := user.Create()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": e.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "注册成功",
	})
}

func ListOne(context *gin.Context) {
	id := context.Param("id")
	user, e := user.Listuser(id)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
	context.AbortWithStatusJSON(200, user)
}

//删除程序，可补充为需要继续输入用户名和密码才能完成的操作
func Cancelletion(context *gin.Context) {
	id := context.Param("id")
	e := user.Deleteuser(id)
	if e != nil {
		// if e.Code == 101 {
		// 	return
		// }
		context.AbortWithStatusJSON(http.StatusBadRequest, e.Error())
	}
}

// 更新用户的手机号信息
func UpdatePhone(context *gin.Context) {
	var num user.UpdatePhone
	if err := context.ShouldBindJSON(&num); err != nil {
		_, e := strconv.Atoi(num.Phone)
		if e != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "电话号码错误！",
			})
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"提示": "参数错误",
		})
	}
	fmt.Println(num)
	e := user.AddPhone(num.ID, num.Phone)
	context.AbortWithStatusJSON(e.Code, e.Error())

}
