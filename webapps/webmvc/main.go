package main

import (
	"html/template"
	"net/http"
	"github.com/gorilla/mux"

	"ogani.com/webmvc/controllers/contact"
	"ogani.com/webmvc/controllers/home"
)
var tmpls *template.Template

func init(){
	tmpls = template.Must(template.ParseFiles("/views"))
}

func addTmpls2HandlerMiddleware(f func(w http.ResponseWriter,r *http.Request, tmpls *template.Template)) http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		f(w,r,tmpls)
	}
}

func main(){
	r:= mux.NewRouter()
	r.HandleFunc("/",addTmpls2HandlerMiddleware(homeController.Index))
	r.HandleFunc("/contact",addTmpls2HandlerMiddleware(contactController.Index)).Methods("GET")
	r.HandleFunc("/contact",addTmpls2HandlerMiddleware(contactController.ContactSubmit)).Methods("POST")
	http.ListenAndServe(":8080",r)
}
