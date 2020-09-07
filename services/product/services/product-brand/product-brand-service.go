package product_brand

import (
	"github.com/jinzhu/gorm"
	"ogani.com/services/product/models"
	"ogani.com/services/product/utility/db"
)

type ProductBrandService struct {
	dbConnector *gorm.DB
}

func (pbs *ProductBrandService) GetAll() (productBrands []models.ProductBrand){
	pbs.dbConnector = db.ConnectDbHandler()
	defer pbs.dbConnector.Close()

	pbs.dbConnector.Model(&models.ProductBrand{}).Find(productBrands)
	return
}