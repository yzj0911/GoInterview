package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello", html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8888", nil))
}
