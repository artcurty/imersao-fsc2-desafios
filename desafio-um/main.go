package main

import (
	"log"
	"net/http"
	"html/template"
)

type Result struct {
	Status string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, Result{})
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Executing server [PORT: 8000]")
	http.ListenAndServe(":8000", nil)
}
