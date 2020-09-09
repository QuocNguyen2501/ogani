package product_type

import (
	"github.com/gofiber/fiber"
	ProductTypeService "ogani.com/services/product/services/product-type"
)

type ProductTypeHandler struct {
	productTypeService *ProductTypeService.ProductTypeService
}

func CreateProductTypeHandler() *ProductTypeHandler{
	return &ProductTypeHandler{
		productTypeService: &ProductTypeService.ProductTypeService{},
	}
}

func (pth *ProductTypeHandler) ProductTypes(c *fiber.Ctx) {
	productTypes := pth.productTypeService.GetAll()

	type ProductTypeDTO struct {
		Id uint `json:"id"`
		Type string `json:"type"`
	}

	var productTypeDTOs []ProductTypeDTO
	for _,item := range productTypes{
		productTypeDTOs = append(productTypeDTOs, ProductTypeDTO{
			Id: item.ID,
			Type: item.Type,
		})
	}

	c.JSON(productTypeDTOs)
}