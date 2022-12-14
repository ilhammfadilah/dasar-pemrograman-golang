package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {
	// var tmpl, err = template.ParseGlob("templates/*")
	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "ilham muhamad fadilah"}
		var tmpl = template.Must(template.ParseFiles(
			"templates/index.html",
			"templates/_header.html",
			"templates/_message.html",
		))
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "about page"}
		var tmpl = template.Must(template.ParseFiles(
			"templates/about.html",
			"templates/_header.html",
			"templates/_message.html",
		))
		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:3000")
	http.ListenAndServe(":3000", nil)
}
