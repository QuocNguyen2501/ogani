package models

import "github.com/jinzhu/gorm"

type ProductType struct{
	gorm.Model
	Type string `gorm:"size:125;not null"`
}