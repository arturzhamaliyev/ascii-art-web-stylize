package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"a-art-w/internal/ascii_art"
)

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/ascii-art", asciiHandler)
	http.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))

	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text, ok := r.Form["input"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	font := r.Form["font"]
	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ascii, err := ascii_art.AsciiArt(text, font)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	templ, err := template.ParseFiles("ui/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, ascii)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if request.Method != http.MethodGet {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("ui/index.html")
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = templ.Execute(writer, nil)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
