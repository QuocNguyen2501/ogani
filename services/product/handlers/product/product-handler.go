package productHandler

import (
	"github.com/gofiber/fiber"
	"github.com/spf13/cast"
	"log"
	"ogani.com/services/product/handlers/product/dto-models"
	ProductService "ogani.com/services/product/services/product"
	ProductBrandService "ogani.com/services/product/services/product-brand"
	ProductTypeService "ogani.com/services/product/services/product-type"
	PagingModel "ogani.com/services/product/utility/models"
	"strconv"
)


type ProductHandler struct {
	productService      *ProductService.ProductService
	productTypeService  *ProductTypeService.ProductTypeService
	productBrandService *ProductBrandService.ProductBrandService
}

func CreateProductHandler() *ProductHandler{
	return &ProductHandler{
		productService : &ProductService.ProductService{},
		productTypeService: &ProductTypeService.ProductTypeService{},
		productBrandService: &ProductBrandService.ProductBrandService{},
	}
}

// Items godoc
// @Summary Get all catalogs
// @Description Get all catalogs
// @Accept json
// @Produce json
// @Param  pageSize query int true "it's page size"
// @Param  pageIndex query int true "it's page index"
// @Success 200
// @BadRequest 400
// @Router /product/items [get]
func (ph *ProductHandler) Items(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	items,totalItems := ph.productService.ProductsWithPaging(pageSize,pageIndex)

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range items{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	responseData := PagingModel.PagingResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		ItemsOnPage: productItemDTOModels,
	}

	c.JSON(responseData)
}

func (ph *ProductHandler) ItemById(c *fiber.Ctx) {
	id, err :=  strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Next(err)
		return
	}

	product := ph.productService.ProductById(id)

	var productDTO dto_models.ProductItemDTOModel
	productDTO.ConvertFromProductItem2DTOModel(&product)

	c.JSON(productDTO)
}

func (ph *ProductHandler) ItemsWithName(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	products, totalItems := ph.productService.ProductsWithName(c.Params("name"),pageSize, pageIndex)

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range products {
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	c.JSON(PagingModel.PagingResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		ItemsOnPage: productItemDTOModels,
	})
}

// Items godoc
// @Summary Get all catalogs
// @Description Get all catalogs by type and brand
// @Accept json
// @Produce json
// @Param  pageSize query int true "it's page size"
// @Param  pageIndex query int true "it's page index"
// @Success 200
// @BadRequest 400
// @Router /product/items/type/{catalogTypeId}/brand/{catalogBrandId}
func (ph *ProductHandler) ItemsByTypeIdAndBrandId(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	products, totalItems := ph.productService.ProductsByTypeIdAndBrandId(cast.ToInt(c.Params("catalogTypeId")),cast.ToInt(c.Params("catalogBrandId")),pageSize, pageIndex)

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range products{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	c.JSON(PagingModel.PagingResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		ItemsOnPage: productItemDTOModels,
	})
}

func (ph *ProductHandler) ItemsByBrandId(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	products,totalItems := ph.productService.ProductsByBrandId(cast.ToInt(c.Params("catalogBrandId")),pageSize, pageIndex)

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range products{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	c.JSON(PagingModel.PagingResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		ItemsOnPage: productItemDTOModels,
	})
}

func (ph *ProductHandler) ProductTypes(c *fiber.Ctx) {
	productTypes := ph.productTypeService.GetAll()

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

func (ph *ProductHandler) ProductBrands(c *fiber.Ctx) {
	productBrands := ph.productBrandService.GetAll()

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

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) {

}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) {

}

func (ph *ProductHandler) DeleteProduct(c *fiber.Ctx) {

}




func getPageSizePageIndex(c *fiber.Ctx) (pageSize, pageIndex int) {
	pageSize,err := strconv.Atoi(c.Query("pageSize","10"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

	pageIndex, err = strconv.Atoi(c.Query("pageIndex","0"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}
	return
}