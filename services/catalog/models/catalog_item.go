package models

import (
	"errors"
	"math"
)

type CatalogItem struct {
	Id int
	Name string
	Description string
	Price float32
	PictureFileName string
	PictureUri string
	CatalogTypeId int
	CatalogType CatalogType
	CatalogBrand CatalogBrand
	AvailableStock int
	RestockThreshold int
	MaxStockThreshold int
	OnReorder bool
}

func (ci CatalogItem) RemoveStock(quantityDesired int) (int,error){
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

func (ci CatalogItem) AddStock(quantity int) int{
	original := ci.AvailableStock
	if (ci.AvailableStock+quantity) > ci.MaxStockThreshold {
		ci.AvailableStock += (ci.MaxStockThreshold - ci.AvailableStock)
	}else{
		ci.AvailableStock += quantity
	}
	ci.OnReorder = false
	return ci.AvailableStock - original
}