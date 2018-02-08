package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var toLove string
		if req.URL.Path[1:] != "" {
			toLove = req.URL.Path[1:]
		} else {
			toLove = "Golang"
		}

		fmt.Fprintf(res, "Hi there, I love %s!", toLove)
	})
	http.ListenAndServe(":8080", nil)
}
