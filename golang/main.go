package main

import (
	"html/template"
	"net/http"
)


type DataPage struct {
	Content string
}

func greeting() string {
	return "Code.education Rocks!"
}

func main() {
	
	renderHtml := template.Must(template.ParseFiles("index.html"))
	content := greeting()
			
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := DataPage{
            Content: content,
        }
		renderHtml.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}