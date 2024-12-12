package main

import (
	"fmt"
	"net/http"
)

type headers int

func (h headers) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mayur-key", "This is header")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Mayur-Headers", "This is how the header sets in code.")
	fmt.Fprint(w, "<h1>Heder is here<h1>")
}

func main() {
	var d headers
	http.ListenAndServe(":8080", d)
}
