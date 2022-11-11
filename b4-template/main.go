package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filePath = path.Join("template", "index.html")
		var tmpl, err = template.ParseFiles(filePath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var data = map[string]interface{}{
			"title": "belajar pemrograman golang",
			"name":  "ilham muhamad fadilah",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:3000")
	http.ListenAndServe(":3000", nil)
}
