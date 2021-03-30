package mapping

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string
	Reviews  []Review `gorm:"foreignKey:AuthorID"`
}

type Product struct {
	gorm.Model
	Upc     string `gorm:"unique"`
	Name    string
	Price   int
	Reviews []Review
}

type Review struct {
	gorm.Model
	Body      string
	AuthorID  uint
	ProductID uint
}
