package user

import (
	"database/sql"
	"time"

	"jinghaijun.com/store/db"
	"jinghaijun.com/store/utils"
	error "jinghaijun.com/store/utils/errors"
)

type UserKind int

const (
	UserKindNormal     UserKind = 1                   //普通用户
	USeUserKindManager UserKind = UserKindNormal << 1 //管理员
)

type User struct {
	ID       uint `gorm:"primarykey"`
	Nickname sql.NullString
	Username string `gorm:"index"` //index根据参数创建索引，多个字段使用相同的名称则创建复合索引
	Password string
	Phone    string `gorm:"index"`
	Gender   uint8
	Avatar   sql.NullString
	Status   uint `gorm:"index"`
	CreateAt time.Time
}

func (user *User) Encrypt_Password() {
	user.Password = utils.Encrypt(user.Password) //加密语句
}

//检查传入数据是否正确，不正确就返回错误
func (user *User) Validate() *error.Error {
	if user.Username == "" || user.Password == "" {
		e := error.New("参数错误")
		e.Code = 400
		return e
	}
	return nil
}

func (user *User) DoesSameUsernameExist() bool {
	var check int64
	db := db.Get_DB()
	db.Exec("select id from users where username = ?", user.Username)
}
