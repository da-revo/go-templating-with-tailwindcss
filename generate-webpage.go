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

func generateWebpage(w http.ResponseWriter, r *http.Request) error {
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		return err
	}
	data := itemPageData{
		PageTitle: "Manipal",
		Items: []item{
			{Title: "Visit EOTT", Done: false},
			{Title: "Swim at the lake", Done: true},
			{Title: "Study?", Done: true},
		},
	}

	tmpl.Execute(w, data)

	return nil
}
