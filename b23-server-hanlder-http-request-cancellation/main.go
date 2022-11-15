package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// do something
	done := make(chan bool)
	go func() {
		// do the process here
		// simulate a long-time request

		// body, err := ioutil.ReadAll(r.Body)

		time.Sleep(10 * time.Second)

		done <- true
	}()

	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("request cancelled")
			} else {
				log.Println("unknow error occured.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
