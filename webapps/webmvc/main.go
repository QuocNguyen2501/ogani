package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"path/filepath"

	"ogani.com/webmvc/controllers/contact"
	"ogani.com/webmvc/controllers/home"
)
var files []string
func init(){
	files = layoutFiles()

}

func addTmpls2HandlerMiddleware(f func(w http.ResponseWriter,r *http.Request, tmplsPath []string)) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		f(w,r,files)
	}
}

func layoutFiles() []string{
	files,err:= filepath.Glob("views/*.html")
	if err != nil {
		panic(err)
	}
	return files
}

func main(){
	fs := http.FileServer(http.Dir("assets/"))

	r:= mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/",addTmpls2HandlerMiddleware(homeController.Index))
	r.HandleFunc("/contact",addTmpls2HandlerMiddleware(contactController.Index)).Methods("GET")

	//r.HandleFunc("/contact",addTmpls2HandlerMiddleware(contactController.ContactSubmit)).Methods("POST")
	http.ListenAndServe(":8080",r)
}
