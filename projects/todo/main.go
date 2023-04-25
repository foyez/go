package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

var tmpl *template.Template

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Todo List",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
		},
	}

	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
