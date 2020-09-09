package product_brand

import (
	"github.com/gofiber/fiber"
	ProductBrandService "ogani.com/services/product/services/product-brand"
)

type ProductBrandHandler struct {
	productBrandService *ProductBrandService.ProductBrandService
}

func CreateProductBrandHandler() *ProductBrandHandler{
	return &ProductBrandHandler{
		productBrandService: &ProductBrandService.ProductBrandService{},
	}
}


func (pbh *ProductBrandHandler) ProductBrands(c *fiber.Ctx) {
	productBrands := pbh.productBrandService.GetAll()

	type ProductBrandDTO struct {
		Id uint `json:"id"`
		Brand string `json:"brand"`
	}
	var productBrandDTOs []ProductBrandDTO
	for _,item := range productBrands {
		productBrandDTOs = append(productBrandDTOs,ProductBrandDTO{
			Id: item.ID,
			Brand: item.Brand,
		})
	}
	c.JSON(productBrandDTOs)
}