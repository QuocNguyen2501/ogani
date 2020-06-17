package contactController

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter,r *http.Request, tmplsPath []string){

	files := append(tmplsPath, "views/contact/contact.html")
	tmpls, _ := template.ParseFiles(files...)
	tmpls.ExecuteTemplate(w,"masterpage",nil)
}

func ContactSubmit(w http.ResponseWriter,r *http.Request, tmplsPath []string){

}