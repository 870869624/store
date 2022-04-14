package authentication

import (
	"errors"

	"jinghaijun.com/store/db"
	"jinghaijun.com/store/models/token"
	"jinghaijun.com/store/models/user"
)

type Authentication struct {
	Username string `gorm:"index"`
	Password string
	ID       uint `gorm:"primarykey"`
}

//验证用户信息，生成token
func (a *Authentication) Create_JWT() (string, error) {
	return token.New(uint64(a.ID), user.UserKindNormal)
}

//登录函数 检测用户信息是否完整，然后检测是否存在，再加密密码对比数据库信息最后登陆成功
func (a *Authentication) SignIn() (string, error) {
	u := &user.User{
		ID:       a.ID,
		Username: a.Username,
		Password: a.Password,
	}
	if err := u.Validate(); err != nil {
		return "", err
	}
	if !u.DoesSameUsernameExist() {
		return "", errors.New("用户不存在")
	}
	u.Encrypt_Password()
	db := db.Get_DB()
	r := db.Table("users").Where("username = ? and password = ?", u.Username, u.Password).First(a)
	if r.Error != nil {
		return "", r.Error
	}
	// r := db.Exec("select * from users where username = ? and password = ? order by id limit 1", u.Username, u.Password)
	// if r.Error != nil {
	// 	return "", r.Error
	// }
	return a.Create_JWT()
}
