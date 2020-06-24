package models

import "github.com/jinzhu/gorm"

type ProductBrand struct{
	gorm.Model
	Id int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Brand string `gorm:"size:125;not null"`
}
