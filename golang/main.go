package main

import (
	"html/template"
	"net/http"
)


Type DataPage struct {
	PageTitle string
}

func greeting() string {
	return ""
}

func main() {
	renderHtml := template.Must(template.ParseFiles("main/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := DataPage{
			PageTitle: greeting()
		}

		renderHtml.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}