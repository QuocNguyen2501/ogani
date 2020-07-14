package models

import "github.com/jinzhu/gorm"

type ProductBrand struct{
	gorm.Model
	Brand string `gorm:"size:125;not null"`
}
