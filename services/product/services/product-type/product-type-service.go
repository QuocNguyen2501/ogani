package product_type

import (
	"github.com/jinzhu/gorm"
	"ogani.com/services/product/models"
	"ogani.com/services/product/utility/db"
)

type ProductTypeService struct{
	dbConnector *gorm.DB
}

func (pts *ProductTypeService) GetAll() (productTypes []models.ProductType){
	pts.dbConnector = db.ConnectDbHandler()
	defer pts.dbConnector.Close()

	pts.dbConnector.Model(&models.ProductType{}).Find(&productTypes)
	return
}