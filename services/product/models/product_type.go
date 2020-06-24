package models

import "github.com/jinzhu/gorm"

type ProductType struct{
	gorm.Model
	Id int `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Type string `gorm:"size:125;not null"`
}