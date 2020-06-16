package handlers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTP(resw http.ResponseWriter, req *http.Request) {
	io.WriteString(resw, "Hello from a HandleFunc #1!\n")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(resw, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resw, "Hello %s\n", body)
}
