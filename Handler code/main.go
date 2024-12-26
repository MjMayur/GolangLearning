package main

import (
	// "fmt"
	"log"
	"net/http"
	"text/template"
)

// type bottle int

// func (b bottle) ServeHTTP(g http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(g, "handle the request here")
// }

// func main() {
// 	var d bottle
// 	http.ListenAndServe(":8080", d)
// }

type bottle int

func (b bottle) ServeHTTP(g http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(g, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d bottle
	http.ListenAndServe(":8080", d)
}
