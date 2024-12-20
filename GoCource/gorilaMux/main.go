package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", param)
	http.HandleFunc("/url", paramWithURL)
	http.HandleFunc("/file", readFile)
	http.ListenAndServe(":8080", nil)
}

// func param(res http.ResponseWriter, req *http.Request) {
// 	v := req.FormValue("q")
// 	io.WriteString(res, "this is in param:"+v)
// }

// Post method sends data in body
func param(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method ="post">
	<input type="text" name="q">
	<input type="Submit">
	</form>
	`+v)
}

// Get method sends the data form url
func paramWithURL(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method ="get">
	<input type="text" name="q">
	<input type="Submit">
	</form>
	`+v)
}

func readFile(res http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Print(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile", f, "\nheader:", h, "\nerr", err)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<form method ="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="Submit">
	</form>
	`+s)
}
