package main

import (
	"io"
	"net/http"
)

type handle1 int

func (h1 handle1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi this is handle1")
}

type handle2 int

func (h2 handle2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi this is handle2")
}
func main() {
	var h1 handle1
	var h2 handle2

	mux := http.NewServeMux()
	//if add for slash / to the rout like ("/handle1/") then it will call on rout like this handle1/something else
	//if don't add for lash / to the rout like ("/handle1/") then it will not going to call on rout like this handle1/something else
	mux.Handle("/handle1", h1)
	mux.Handle("/handle2", h2)

	//#############--------------another approach to to same thing is------------#####
	// var h1 handle1
	// var h2 handle2

	// http.Handle("/handle1", h1)
	// http.Handle("/handle2", h2)
	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", mux)
}

// ############-------------approach 2--------------############

// func h1 ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "Hi this is handle1")
// }

// func h2 ServeHTTP(res http.ResponseWriter, req *http.Request) {
// 	io.WriteString(res, "Hi this is handle2")
// }

// func main() {

// 	http.Handle("/handle1", h1)
// 	http.Handle("/handle2", h2)

// 	http.ListenAndServe(":8080", nil)
// }
