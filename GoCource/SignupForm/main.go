package main

import "net/http"

func main() {
	http.HandleFunc("/", signup)
	http.HandleFunc("/bar", handleBar)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}
