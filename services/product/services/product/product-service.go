package product

import (
	"github.com/jinzhu/gorm"
	"ogani.com/services/product/models"
	"ogani.com/services/product/utility/db"
)

type ProductService struct{
	dbConnector *gorm.DB
}


func (ps *ProductService) ProductsWithPaging(pageSize, pageIndex int) (products []models.ProductItem, totalItems int){
	ps.dbConnector = db.ConnectDbHandler()
	defer ps.dbConnector.Close()

	totalItemsCh := make(chan int)
	go func(){
		var _totalItems = 0
		ps.dbConnector.Model(&models.ProductItem{}).Count(&_totalItems)
		totalItemsCh <- _totalItems
	}()
	itemsOnPageCh := make(chan []models.ProductItem)
	go func() {
		var _itemsOnPage []models.ProductItem
		ps.dbConnector.Model(&models.ProductItem{}).Order("name asc").Offset(pageIndex * pageSize).Limit(pageSize).Find(&_itemsOnPage)
		itemsOnPageCh <- _itemsOnPage
	}()

	totalItems = <-totalItemsCh
	products = <- itemsOnPageCh
	return
}

func (ps *ProductService) ProductById(id int) (product models.ProductItem) {
	ps.dbConnector = db.ConnectDbHandler()
	defer ps.dbConnector.Close()

	ps.dbConnector.Model(&models.ProductItem{}).Where("id = ?", id).Find(&product)

	product.FillProductUrl("base url")
	return
}

func (ps *ProductService) ProductsWithName(name string, pageSize, pageIndex int) (products []models.ProductItem, totalItems int){
	ps.dbConnector = db.ConnectDbHandler()
	defer ps.dbConnector.Close()

	query := ps.dbConnector.Model(&models.ProductItem{}).Where("name LIKE %?%", name)
	totalChan := make(chan int,1)
	go func() {
		totalItems := 0
		query.Count(&totalItems)
		totalChan <- totalItems
	}()

	itemsChan := make(chan []models.ProductItem)
	go func(){
		var itemsOnPage []models.ProductItem
		query.Order("name asc").Offset(pageIndex * pageSize).Limit(pageSize).Find(&itemsOnPage)
		itemsChan <- itemsOnPage
	}()

	products = <- itemsChan
	totalItems = <- totalChan

	return
}

func (ps *ProductService) ProductsByTypeIdAndBrandId(catalogTypeId,catalogBrandId,pageSize, pageIndex int) (products []models.ProductItem, totalItems int){
	ps.dbConnector = db.ConnectDbHandler()
	defer ps.dbConnector.Close()

	query:= ps.dbConnector.Model(&models.ProductItem{})
	if catalogTypeId != 0{
		query.Where("catalogTypeId = ?", catalogTypeId)
	}

	if catalogBrandId != 0 {
		query.Where("catalogBrandId = ?", catalogBrandId)
	}

	totalItemsChan := make(chan int)
	go func(){
		var totalItems int
		query.Count(&totalItems)
		totalItemsChan <- totalItems
	}()

	itemsOnPageCh := make(chan []models.ProductItem)

	go func(){
		var itemsOnPage []models.ProductItem
		query.Offset(pageIndex * pageSize).Limit(pageSize).Find(&itemsOnPage)
		itemsOnPageCh <- itemsOnPage
	}()
	products = <- itemsOnPageCh
	totalItems = <- totalItemsChan
	return
}

func (ps *ProductService) ProductsByBrandId(catalogBrandId,pageSize, pageIndex int) (products []models.ProductItem, totalItems int) {
	ps.dbConnector = db.ConnectDbHandler()
	defer ps.dbConnector.Close()

	query:= ps.dbConnector.Model(&models.ProductItem{})
	if catalogBrandId != 0 {
		query.Where("catalogBrandId = ?", catalogBrandId)
	}

	totalItemsChan := make(chan int)
	go func(){
		var totalItems int
		query.Count(&totalItems)
		totalItemsChan <- totalItems
	}()

	itemsOnPageCh := make(chan []models.ProductItem)

	go func(){
		var itemsOnPage []models.ProductItem
		query.Offset(pageIndex * pageSize).Limit(pageSize).Find(&itemsOnPage)
		itemsOnPageCh <- itemsOnPage
	}()

	products = <-itemsOnPageCh
	totalItems = <- totalItemsChan
	return
}