package productController

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"ogani.com/services/product/models"
	"strconv"
)

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
	pageSize,err := strconv.Atoi(c.Query("pageSize","10"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("pageIndex","0"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

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
	itemsOnPage := <- itemsOnPageCh

	c.Send(struct {
		pageIndex int
		pageSize int
		totalItems int
		itemsOnPage []models.ProductItem
	}{
		pageIndex: pageIndex,
		pageSize: pageSize,
		totalItems: totalItems,
		itemsOnPage: itemsOnPage,
	})
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

	c.Send(product)
}

func ItemsWithName(c *fiber.Ctx) {
	pageSize,err := strconv.Atoi(c.Query("pageSize","10"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("pageIndex","0"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

	db:= connectDbHandler()

	totalItems := 0
	db.Model(&models.ProductItem{}).Where("name LIKE %?%",c.Params("name")).Count(&totalItems)

	var itemsOnPage []models.ProductItem
	db.Model(&models.ProductItem{}).Where("name LIKE %?%",c.Params("name")).Order("name asc").Offset(pageIndex * pageSize).Limit(pageSize).Find(&itemsOnPage)

	c.Send(struct {
		pageIndex int
		pageSize int
		totalItems int
		itemsOnPage []models.ProductItem
	}{
		pageIndex: pageIndex,
		pageSize: pageSize,
		totalItems: totalItems,
		itemsOnPage: itemsOnPage,
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
	pageSize,err := strconv.Atoi(c.Query("pageSize","10"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("pageIndex","0"))
	if err !=nil{
		log.Fatalln(err)
		c.Next(err)
		return
	}
	db:= connectDbHandler()

	var itemsOnPage []models.ProductItem
	query:= db.Model(&models.ProductItem{}).Where("catalogTypeId = ?", c.Params("catalogTypeId"))

	if c.Params("catalogBrandId") != "" {
		query.Where("catalogBrandId = ?", c.Params("catalogBrandId"))
	}
	var totalItems int
	query.Count(&totalItems)
	query.Offset(pageIndex * pageSize).Limit(pageSize).Find(&itemsOnPage)
	c.Send(struct {
		pageIndex int
		pageSize int
		totalItems int
		itemsOnPage []models.ProductItem
	}{
		pageIndex: pageIndex,
		pageSize: pageSize,
		totalItems: totalItems,
		itemsOnPage: itemsOnPage,
	})
}

func ItemsByBrandId(c *fiber.Ctx) {

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

type ProductItemDTO struct {
	ID        			uint
	Name              string
	Description       string
	Price             float32
	PictureFileName   string
	PictureUri        string
	ProductTypeId     int
	ProductBrandId     int
	AvailableStock    int
	RestockThreshold  int
	MaxStockThreshold int
	OnReorder         bool
}