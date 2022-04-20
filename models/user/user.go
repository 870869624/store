package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"jinghaijun.com/store/db"
	"jinghaijun.com/store/utils"
	"jinghaijun.com/store/utils/errors"
)

type UserKind int

const (
	UserKindNormal     UserKind = 1                   //普通用户
	USeUserKindManager UserKind = UserKindNormal << 1 //管理员
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Nickname  sql.NullString
	Username  string `gorm:"index"` //index根据参数创建索引，多个字段使用相同的名称则创建复合索引
	Password  string
	Phone     string `gorm:"index"`
	Gender    uint8
	Avatar    sql.NullString
	Status    uint `gorm:"index"`
	CreatedAt time.Time
}

type UpdatePhone struct {
	ID    int
	Phone string
}

func (user *User) Encrypt_Password() {
	user.Password = utils.Encrypt(user.Password) //加密语句
}

//检查传入数据是否正确，不正确就返回错误
func (user *User) Validate() *errors.Error {
	if user.Username == "" || user.Password == "" {
		e := errors.New("参数错误")
		e.Code = 400
		return e
	}
	return nil
}

func (user *User) DoesSameUsernameExist() bool {
	var count int64
	db := db.Get_DB()
	db.Table("users").Where("username = ?", user.Username).Count(&count)
	return count > 0
}

//插入数据库的模块1.检测是否数据完整2.然后检测是否存在3.然后加密写入数据库
func (user *User) Create() *errors.Error {
	db := db.Get_DB()
	if err := user.Validate(); err != nil {
		err.Code = 400
		return err
	}
	if user.DoesSameUsernameExist() {
		e := errors.New("该用户已经存在")
		e.Code = http.StatusBadRequest
		return e
	}
	user.Encrypt_Password()
	e := db.Exec("insert into users (id, nickname, username, password, phone, gender, avatar, status) values (?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Nickname,
		user.Username, user.Password, user.Phone, user.Gender, user.Avatar, user.Status)
	if e.Error != nil {
		result := errors.New(e.Error.Error())
		return result
	}
	return nil

	//result := db.Create(user) 这也是插入语句
}

//根据获取的id查找用户的信息并且返回信息结构体与对应错误
func Listuser(id string) (*User, error) {
	var listuser User
	i, e := strconv.Atoi(id)
	if e != nil {
		return nil, e
	}
	db := db.Get_DB()
	err := db.Table("users").Where("id = ?", i).First(&listuser)
	if err.Error != nil {
		return nil, err.Error
	}
	return &listuser, nil
}

//用户注销(包方法)
func Deleteuser(id string) *errors.Error {
	var deletuser User
	i, e := strconv.Atoi(id)
	if e != nil {
		result := errors.New("传入参数错误！")
		result.Code = 405
		return result
	}
	db := db.Get_DB()
	// r := db.Raw("select * from users where id = ?", i).Scan(&deletuser)
	r := db.Table("users").Where("id = ?", i).First(&deletuser)
	if r.Error != nil {
		c := errors.New("该用户不存在！")
		c.Code = 406
		return c
	}
	fmt.Println(deletuser.Username)
	if deletuser.Phone == "" {
		a := errors.New("请验证手机号！") //令一个a为errors类型
		a.Code = 101
		return a
	}
	db.Exec("delete from users where id = ?", i)
	b := errors.New("注销成功")
	b.Code = 200
	return b
}

//增加手机号到数据库中
// func AddPhone(id int, phone string) *errors.Error {
// 	if len(phone) != 11 {
// 		result := errors.New("电话号码输入错误！")
// 		return result
// 	}
// 	db := db.Get_DB()
// 	db.Exec("update users set phone = ? where id = ?", phone, id)
// 	return nil
// }

//比较信息
// func (u *User, ) contrast() bool {
// 	if
// }

//更新用户信息
func (user *User) Update() error {
	var first User
	fmt.Println(user.Username)
	if user.Username == "" || user.Password == "" {
		result := errors.New("参数错误")
		return result
	}
	if user.Phone == "" {
		result := errors.New("请验证电话号码")
		return result
	}
	if len(user.Phone) != 11 {
		result := errors.New("电话号码错误")
		return result
	}
	user.Encrypt_Password()
	db := db.Get_DB()
	db.Where("id = ?", user.ID).First(&first)
	e := db.Table("users").Where("id = ?", user.ID).First(&first)
	if e.Error != nil {
		return errors.New(e.Error.Error())
	}
	connection := db.Model(first).Updates(user)
	if connection.Error != nil {
		return errors.New(e.Error.Error())
	}
	return nil
}
