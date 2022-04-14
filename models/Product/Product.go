package product

import "time"

type Product struct {
	ID          int    `json:"id" gorm:"paimaryKey"`
	Title       string `json:"title"`
	Price       int
	SubTitle    string    `json:"sub_title"`
	Thumbnail   string    `json:"thumbnail"`
	Pictures    string    `json:"pictures" gorm:"type:JSON"`
	CatalogueId int       `json:"catalogue_id"`
	Sort        uint8     `json:"sort"`
	Status      uint      `json:"status" gorm:"index"`
	CreatedAt   time.Time `json:"created_at"`
}

//检测输入信息完整与否
func (p *Product) productValidate() bool {
	return p.Title != "" && p.SubTitle != "" && p.Thumbnail != "" && p.CatalogueId != 0
}

//新增商品
func (p *Product) Create() error {
	return nil
}
