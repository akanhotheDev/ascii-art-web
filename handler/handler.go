package handler

import (
	"ascii-art-web/asciiart"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("Templates/index.html"))

type pageData struct {
	Result string
	Banner string
	Text   string
	Error  string
}

// GET → show page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != ("/") {
		http.Error(w, "404 Page not found", http.StatusNotFound) // Hanlding Wrong Path
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Hanling Method of form submitting
		return
	}

	data := pageData{}
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError) // 500 Internal Server Error
	}
}

// POST → process form
func AsciiHomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("passed-text")
	bannerStyle := r.FormValue("fonts")
	
	// when text is not provided || Banner Not selected
	if bannerStyle == "" || text == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `
		<style>
		body   {
		background-color: #333;
		}
		</style>
		<center>
		<h1>400</h1>
    	<h2>Bad Request</h2>
		</center>
		`)
			return
	}

	

	result := asciiart.Render(text, "banner/"+bannerStyle+".txt")

	data := pageData{
		Result: result,
		Banner: bannerStyle,
		Text:   text,
	}

	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "error rendering page", http.StatusInternalServerError)
		return
	}
}
