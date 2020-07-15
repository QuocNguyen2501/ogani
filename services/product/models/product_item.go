package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"math"
)

type ProductItem struct {
	gorm.Model
	Name              string `gorm:"size:125;not null"`
	Description       string `gorm:"size:255"`
	Price             float32 `gorm:"not null"`
	PictureFileName   string `gorm:"not null"`
	PictureUri        string `gorm:"not null"`
	ProductTypeId     int `gorm:"not null"`
	ProductType       ProductType `gorm:"not null;association_autoupdate:false;association_autocreate:false;foreignkey:ProductTypeId"`
	ProductBrandId     int `gorm:"not null"`
	ProductBrand      ProductBrand `gorm:"not null;association_autoupdate:false;association_autocreate:false;foreignkey:ProductBrandId"`
	AvailableStock    int `gorm:"not null"`
	RestockThreshold  int `gorm:"not null"`
	MaxStockThreshold int `gorm:"not null"`
	OnReorder         bool `gorm:"not null"`
}

func (ci ProductItem) RemoveStock(quantityDesired int) (int,error){
	if ci.AvailableStock == 0 {
		return 0,errors.New("Empty stock, product item "+ci.Name+" is sold out")
	}
	if quantityDesired <= 0{
		return 0,errors.New("Item units desired should be greater than zero")
	}
	removed := math.Min(float64(quantityDesired),float64(ci.AvailableStock))
	ci.AvailableStock -= int(removed)
	return int(removed),nil
}

func (ci ProductItem) AddStock(quantity int) int{
	original := ci.AvailableStock
	if (ci.AvailableStock+quantity) > ci.MaxStockThreshold {
		ci.AvailableStock += (ci.MaxStockThreshold - ci.AvailableStock)
	}else{
		ci.AvailableStock += quantity
	}
	ci.OnReorder = false
	return ci.AvailableStock - original
}

func (ci ProductItem) FillProductUrl(picBaseUrl string){
	// do something get base url
}