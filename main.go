package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates map[string]*template.Template

func init() {
	loadTemplates()
}

func loadTemplates() {
	var baseTemplate = "templates/layout/_base.html"
	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/index.html"))
}

func main() {
	router := mux.NewRouter()

	ServeStatic(router, "public/")

	router.HandleFunc("/", HomeHandler).Methods("GET")

	log.Printf("Listening on port: %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), router))
}

func ServeStatic(router *mux.Router, staticDirectory string) {
	staticPaths := map[string]string{
		"assets":  staticDirectory + "assets/",
		"styles":  staticDirectory + "styles/",
		"scripts": staticDirectory + "scripts/",
	}
	for pathName, pathValue := range staticPaths {
		pathPrefix := "/" + pathName + "/"
		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
			http.FileServer(http.Dir(pathValue))))
	}
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	if err := templates["index"].Execute(res, nil); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
