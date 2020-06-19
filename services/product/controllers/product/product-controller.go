package productController

import "net/http"

// Items godoc
// @Summary Get all catalogs
// @Description Get all catalogs
// @Accept json
// @Produce json
// @Param  size query int true "it's page size"
// @Param  pageIndex query int true "it's page index"
// @Success 200
// @Router /product/items [get]
func Items(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ItemById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ItemsWithName(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ItemsByTypeIdAndBrandId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ItemsByBrandId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProductTypes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ProductBrands(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}