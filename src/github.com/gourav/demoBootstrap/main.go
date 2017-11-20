package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/views/*.html"))
}

func main() {

	http.Handle("/appbundle_js/", http.StripPrefix("/appbundle_js/", http.FileServer(http.Dir("./dist"))))
	http.Handle("/fabicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.go.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
