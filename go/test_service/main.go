package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test_api", test_handler)
	http.ListenAndServe("0.0.0.0:5000", nil)
}

func test_handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOLA TEST API"))
}
