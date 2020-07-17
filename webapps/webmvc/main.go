package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"path/filepath"

	"ogani.com/webmvc/controllers/blog"
	"ogani.com/webmvc/controllers/blog-details"
	"ogani.com/webmvc/controllers/cart"
	"ogani.com/webmvc/controllers/checkout"
	"ogani.com/webmvc/controllers/contact"
	"ogani.com/webmvc/controllers/home"
	"ogani.com/webmvc/controllers/product-details"
	"ogani.com/webmvc/controllers/products"
)

var files []string

func init() {
	files = layoutFiles()
}

func addTmpls2HandlerMiddleware(f func(w http.ResponseWriter, r *http.Request, tmplsPath []string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, files)
	}
}

func layoutFiles() []string {
	files, err := filepath.Glob("views/*.html")
	if err != nil {
		panic(err)
	}
	return files
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.HandleFunc("/", addTmpls2HandlerMiddleware(homeController.Index)).Methods("GET")
	r.HandleFunc("/products", addTmpls2HandlerMiddleware(productsController.Index)).Methods("GET")
	r.HandleFunc("/product-details/{id}", addTmpls2HandlerMiddleware(productDetailsController.Index)).Methods("GET")
	r.HandleFunc("/cart", addTmpls2HandlerMiddleware(cartController.Index)).Methods("GET")
	r.HandleFunc("/checkout", addTmpls2HandlerMiddleware(checkoutController.Index)).Methods("GET")
	r.HandleFunc("/blog", addTmpls2HandlerMiddleware(blogController.Index)).Methods("GET")
	r.HandleFunc("/blog-details/{id}", addTmpls2HandlerMiddleware(blogDetailsController.Index)).Methods("GET")
	r.HandleFunc("/contact", addTmpls2HandlerMiddleware(contactController.Index)).Methods("GET")
	r.HandleFunc("/contact-submit", addTmpls2HandlerMiddleware(contactController.ContactSubmit)).Methods("POST")

	http.ListenAndServe(":4200", r)
}
