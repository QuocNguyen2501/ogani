package productDetailsController

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter,r *http.Request, tmplsPath []string){

	files := append(tmplsPath, "views/product-details/product-details.html")
	tmpls, _ := template.ParseFiles(files...)
	tmpls.ExecuteTemplate(w,"masterpage",nil)
}
