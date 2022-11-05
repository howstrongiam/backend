package models

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserRefer uint
	ProductId uint
	Product   Product `gorm:"foreignKey:ProductId"`
}
