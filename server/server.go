package main

import (
	"fmt"
	"net/http"
)

func main() {
	const PORT = 8080

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var toLove string = "GoLang"
		if req.URL.Path[1:] != "" {
			toLove = req.URL.Path[1:]
		}

		fmt.Fprintf(res, "Hi there, I love %s!", toLove)
	})

	fmt.Printf("App up at http://localhost:%v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}
