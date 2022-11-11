package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
	fmt.Println("server started at localhost:7000")
	http.ListenAndServe(":7000", nil)
}
