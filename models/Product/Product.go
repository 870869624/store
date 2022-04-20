package product

import (
	"fmt"
	"time"

	"jinghaijun.com/store/db"
	"jinghaijun.com/store/models/catalogue"
	"jinghaijun.com/store/utils/errors"
)

type Product struct {
	ID          int       `json:"id" gorm:"paimaryKey"`
	Title       string    `json:"title"` //产品标题
	Price       int       //产品价格
	SubTitle    string    `json:"sub_title"`                 //产品副标题
	Thumbnail   string    `json:"thumbnail"`                 //产品缩略图
	Pictures    string    `json:"pictures" gorm:"type:JSON"` //产品图片
	CatalogueId int       `json:"catalogue_id"`              //分类ID
	Description string    `json:"description"`               //商品描述
	Sort        uint8     `json:"sort"`                      //顺序
	Status      uint      `json:"status" gorm:"index"`       //状态 0 开启 1 下架 2 售罄
	CreatedAt   time.Time `json:"created_at"`
}

//检测输入信息完整与否
func (p *Product) validate() bool {
	return p.Title != "" && p.SubTitle != "" && p.Thumbnail != "" && p.CatalogueId != 0
}

//新增商品
func (p *Product) check() error {
	if !p.validate() {
		return errors.New("参数错误")
	}
	cata, e := catalogue.One(p.CatalogueId)
	if e != nil || cata.ID == 0 {
		return errors.New("分类不存在!")
	}
	return nil
}
func (p *Product) Create() error {
	e := p.check()
	if e != nil {
		return e
	}
	fmt.Println(p)
	db := db.Get_DB()
	result := db.Create(p)
	return result.Error
}
func (p *Product) Delete(id int) error {
	p.check()
	e := p.check()
	if e != nil {
		return e
	}
	db := db.Get_DB()
	action := db.Table("products").Where("id = ?", id)
	if action.Error != nil {
		return action.Error
	}
	return nil
}
func (p *Product) Update() error {
	var first Product
	if !p.validate() {
		return errors.New("参数错误")
	}
	db := db.Get_DB()
	e := db.Table("products").Where("id = ?", p.ID).First(&first)
	if e.Error != nil {
		return e.Error
	}
	action := db.Model(first).Updates(p)
	return action.Error
}

//包方法
type Pagination struct {
	Page   int `gorm:"page"`   //每页多少条
	Number int `gorm:"number"` //当前多少页
}

type Order struct {
	By []string
}
type ProductsearchQuery struct {
	Pagination
	Order
	Title        string
	Price_gt     int
	Price_lt     int
	Catalogue_ID uint64
}
type ProductListResponse struct {
	Pagination
	Result []Product `json:"result"`
}

//展示商品
func List(query *ProductsearchQuery) (*ProductListResponse, error) {
	reponse := &ProductListResponse{
		Pagination: query.Pagination,
	}
	db := db.Get_DB()
	request := db.Table("products")
	if query.Title != "" {
		request = request.Where("title LIKE ?", fmt.Sprintf("%s%%", query.Title))
	}
	if query.Price_gt != 0 {
		request = request.Where("price >= ?", query.Price_gt)
	}
	if query.Price_lt != 0 {
		request = request.Where("price < ?", query.Price_lt)
	}
	if query.Catalogue_ID != 0 {
		request = request.Where("catalogue_id = ?", query.Catalogue_ID)
	}
	offset := query.Number * (query.Page - 1)
	request = request.Offset(offset).Limit(query.Number)
	for _, v := range query.By {
		request = request.Order(v)
	}
	e := request.Preload("Catalogue").Find(&reponse.Result)
	return reponse, e.Error
}
