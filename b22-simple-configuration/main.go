package main

import (
	"b22-simple-configuration/conf"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CustomMux struct {
	http.ServeMux
}

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if conf.Configuration.Log.Verbose {
		log.Println("incoming request from", r.Host, "accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CustomMux)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("how are you"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = conf.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = conf.Configuration.Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", conf.Configuration().Server.Port)

	if conf.Configuration().Log.Verbose {
		log.Printf("Starting server at %s\n", server.Addr)
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
