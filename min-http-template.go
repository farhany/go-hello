package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("index.html"))
		t.Execute(w, nil)
	})

	fmt.Println("Listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
