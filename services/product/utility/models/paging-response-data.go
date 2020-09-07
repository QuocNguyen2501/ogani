package models

type PagingResponseData struct {
	PageIndex   int         `json:"page_index"`
	PageSize    int         `json:"page_size"`
	TotalItems  int         `json:"total_items"`
	ItemsOnPage interface{} `json:"items_on_page"`
}