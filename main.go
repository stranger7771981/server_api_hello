package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	name := r.URL.Query().Get("name")

	message := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":80", nil)
}
