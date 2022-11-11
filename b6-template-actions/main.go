package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "ilham muhamad fadilah",
			Gender:  "male",
			Hobbies: []string{"reading books", "sleeping", "playing game"},
			Info:    Info{"google", "cijeungjing"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 divisions"
}
