package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/shorten", shorten)

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	l := filepath.Join("static", "template", "layout.html")
	p := filepath.Join("static", "template", "index.html")
	t, err := template.ParseFiles(l, p)
	if err != nil {
		log.Printf("error in parsing template: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w, "layout", nil)
}

func shorten(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Printf("error in parsing form: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		fmt.Fprintf(w, "Bad Request")
		return
	}
	url := req.FormValue("url")
	fmt.Fprintf(w, "URL: %s\n", url)
}
