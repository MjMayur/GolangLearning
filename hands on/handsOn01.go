package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hii this is me mayur")
}

func b(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "How are you")
}

func main() {
	http.HandleFunc("/", d)
	http.HandleFunc("/dog", b)
	http.HandleFunc("/me", c)
	http.ListenAndServe(":8080", nil)
}
