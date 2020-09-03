package dto_models

import (
	"ogani.com/services/product/models"
)

type ProductItemDTOModel struct{
	ID        			uint `json:"id"`
	Name      			string `json:"name"`
	Description       	string `json:"description"`
	Price             	float32 `json:"price"`
	PictureFileName   	string	`json:"picture_file_name"`
	PictureUri 			string `json:"picture_uri"`
	ProductTypeId 		int `json:"product_type_id"`
	ProductType 		string `json:"product_type"`
	ProductBrandId 		int `json:"product_brand_id"`
	ProductBrand 		string `json:"product_brand"`
	AvailableStock    	int `json:"available_stock"`
	RestockThreshold  	int `json:"restock_threshold"`
	MaxStockThreshold 	int `json:"max_stock_threshold"`
	OnReorder         	bool `json:"on_reorder"`
}

func (p *ProductItemDTOModel) ConvertFromProductItem2DTOModel(pi *models.ProductItem){
	p.ID = pi.ID
	p.Name = pi.Name
	p.Description = pi.Description
	p.Price = pi.Price
	p.PictureFileName = pi.PictureFileName
	p.PictureUri = pi.PictureUri
	p.ProductTypeId = pi.ProductTypeId
	p.ProductType = pi.ProductType.Type
	p.ProductBrandId = pi.ProductBrandId
	p.ProductBrand = pi.ProductBrand.Brand
	p.AvailableStock = pi.AvailableStock
	p.RestockThreshold = pi.RestockThreshold
	p.MaxStockThreshold = pi.MaxStockThreshold
	p.OnReorder = pi.OnReorder
}