package productController

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"ogani.com/services/product/controllers/product/dto-models"
	"ogani.com/services/product/models"
	"strconv"
)

type ResponseData struct {
	PageIndex   int         `json:"page_index"`
	PageSize    int         `json:"page_size"`
	TotalItems  int         `json:"total_items"`
	ItemsOnPage interface{} `json:"items_on_page"`
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
func Items(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	db := connectDbHandler()
	defer db.Close()

	totalItemsCh := make(chan int)
	go func(){
		var _totalItems = 0
		db.Model(&models.ProductItem{}).Count(&_totalItems)
		totalItemsCh <- _totalItems
	}()
	itemsOnPageCh := make(chan []models.ProductItem)
	go func() {
		var _itemsOnPage []models.ProductItem
		db.Model(&models.ProductItem{}).Order("name asc").Offset(pageIndex * pageSize).Limit(pageSize).Find(&_itemsOnPage)
		itemsOnPageCh <- _itemsOnPage
	}()

	totalItems := <-totalItemsCh

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range <- itemsOnPageCh{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	responseData := ResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  totalItems,
		ItemsOnPage: productItemDTOModels,
	}

	c.JSON(responseData)
}

func ItemById(c *fiber.Ctx) {
	id, err :=  strconv.Atoi(c.Params("id"))
	if err != nil {
		c.Next(err)
		return
	}

	db := connectDbHandler()
	defer db.Close()

	var product models.ProductItem
	db.Model(&models.ProductItem{}).Where("id = ?", id).Find(&product)

	product.FillProductUrl("base url")

	var productDTO dto_models.ProductItemDTOModel
	productDTO.ConvertFromProductItem2DTOModel(&product)

	c.JSON(productDTO)
}

func ItemsWithName(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	db:= connectDbHandler()
	defer db.Close()

	query := db.Model(&models.ProductItem{}).Where("name LIKE %?%", c.Params("name"))
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

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range <-itemsChan {
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}


	c.JSON(ResponseData {
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  <-totalChan,
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
func ItemsByTypeIdAndBrandId(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	db:= connectDbHandler()
	defer db.Close()

	query:= db.Model(&models.ProductItem{})
	if c.Params("catalogTypeId") != ""{
		query.Where("catalogTypeId = ?", c.Params("catalogTypeId"))
	}

	if c.Params("catalogBrandId") != "" {
		query.Where("catalogBrandId = ?", c.Params("catalogBrandId"))
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

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range <- itemsOnPageCh{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	c.JSON(ResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  <- totalItemsChan,
		ItemsOnPage: productItemDTOModels,
	})
}

func ItemsByBrandId(c *fiber.Ctx) {
	pageSize, pageIndex := getPageSizePageIndex(c)

	db:= connectDbHandler()
	defer db.Close()

	query:= db.Model(&models.ProductItem{})
	if c.Params("catalogBrandId") != "" {
		query.Where("catalogBrandId = ?", c.Params("catalogBrandId"))
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

	var productItemDTOModels []dto_models.ProductItemDTOModel
	for _,item := range <- itemsOnPageCh{
		var dto  dto_models.ProductItemDTOModel
		dto.ConvertFromProductItem2DTOModel(&item)
		productItemDTOModels = append(productItemDTOModels,dto)
	}

	c.JSON(ResponseData{
		PageIndex:   pageIndex,
		PageSize:    pageSize,
		TotalItems:  <- totalItemsChan,
		ItemsOnPage: productItemDTOModels,
	})
}

func ProductTypes(c *fiber.Ctx) {

}

func ProductBrands(c *fiber.Ctx) {

}

func UpdateProduct(c *fiber.Ctx) {

}

func CreateProduct(c *fiber.Ctx) {

}

func DeleteProduct(c *fiber.Ctx) {

}


func connectDbHandler() *gorm.DB{
	db, err := gorm.Open(cast.ToString(viper.Get("dbDialect")), viper.Get(cast.ToString("connStr")))
	if err != nil{
		panic(err)
	}

	return db
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