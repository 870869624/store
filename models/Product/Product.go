package product

type Product struct {
	ID        int `gorm:"paimaryKey"`
	Title     string
	Price     int
	SubTitle  string
	Thumbnail string
	Pictures  string
}
