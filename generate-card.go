package main

import (
	"net/http"
	"text/template"
)

type itemPageData struct {
	PageTitle string
	Items     []item
}

type item struct {
	Title string
	Done  bool
}

func generateCard(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		return err
	}
	data := itemPageData{
		PageTitle: "My TODO list",
		Items: []item{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl.Execute(w, data)

	return nil
}
