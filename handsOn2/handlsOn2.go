package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hii this is me mayur")
}

func b(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal("Error occurred", err)
	}
	err = tpl.ExecuteTemplate(res, "index.gohtml", "Mayur")
	if err != nil {
		log.Fatal("Error is here", err)
	}
}

func main() {
	http.HandleFunc("/", d)
	http.HandleFunc("/dog", b)
	http.HandleFunc("/me", c)
	http.ListenAndServe(":8080", nil)
}
