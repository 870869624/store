package catalogue

import (
	"errors"
	"fmt"
	"time"

	"jinghaijun.com/store/db"
)

type Catalogue struct {
	ID        int       `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"column:TITLE"` //分类标题
	Thumbnail string    `json:"thumbnail"`                 //分类图标
	Sort      uint8     `json:"sort"`                      //排序
	Status    uint      `json:"status" gorm:"index"`       //状态 0 开启 1 禁用
	CreatedAt time.Time `json:"created_at"`
}

//检查信息是否完整
func (c *Catalogue) validate() bool {
	fmt.Println(c, c.Title, c.Thumbnail)
	return c.Title != "" && c.Thumbnail != ""
}

//检测是否已经存在
func (c *Catalogue) doesNameExit() bool {
	var count int64
	db := db.Get_DB()
	db.Table("catalogues").Where("id = ?", c.ID).Count(&count)
	return count > 0
}

//创建分类
func (c *Catalogue) Creat() error {
	if !c.validate() {
		return errors.New("参数错误")
	}
	if c.doesNameExit() {
		return errors.New("该分类已经存在")
	}
	db := db.Get_DB()
	result := db.Create(c)
	return result.Error
}

type Pagination struct {
	Page   int `gorm:"page"`
	Number int `gorm:"number"`
}
type Order struct {
	By []string
}
type CatalogueQuery struct {
	Pagination
	Order
	Title string
}
type CatalogueResponse struct {
	Pagination
	Response []Catalogue `gorm:"response"`
}

//包方法 检测某个产品的分类
func One(id int) (*Catalogue, error) {
	var one Catalogue
	db := db.Get_DB()
	r := db.First(&one, id)
	return &one, r.Error
}
func Listall(query *CatalogueQuery) (*CatalogueResponse, error) {
	response := &CatalogueResponse{
		Pagination: query.Pagination,
	}
	db := db.Get_DB()
	request := db.Table("CATALOGUES")
	if query.Title != "" {
		request = request.Where("title like ?", fmt.Sprintf("%s%%", query.Title))
	}
	offset := query.Number * (query.Page - 1)
	request = request.Offset(offset).Limit(query.Number)
	for _, v := range query.By {
		request = request.Order(v)
	}
	e := request.Find(&response.Response)
	return response, e.Error
}
func ListOne(query *CatalogueQuery) (*CatalogueResponse, error) {
	response := &CatalogueResponse{
		Pagination: query.Pagination,
	}
	db := db.Get_DB()
	req := db.Table("CATALOGUES")
	if query.Title != "" {
		req = req.Where("title like ?", fmt.Sprintf("%s%%", query.Title))
	}
	result := req.Find(&response.Response)
	return response, result.Error
}
func (c *Catalogue) Update() error {
	var update Catalogue
	db := db.Get_DB()
	connect := db.Table("catalogues").Where("id = ?", c.ID).First(&update)
	if connect.Error != nil {
		return connect.Error
	}
	result := db.Model(update).Updates(c)
	if result.Error != nil {
		return result.Error
	}
	return errors.New("更新完成")
}
