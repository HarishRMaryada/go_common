package main

import (
	model "go_common/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
		body, _ := ioutil.ReadAll(r.Body)
		log.Printf("Body %s\n", body)
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	model.Model()
	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	http.ListenAndServe(":8080", nil)
}
