package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/qbradq/petshop-go/data"
)

// tpl is the pre-parsed template file system
var tpl *template.Template

// templateHandler renders HTML templates.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rn := path.Base(r.RequestURI)
	tn := "index.html.tpl"
	if rn != "/" && rn != "." {
		tn = rn + ".tpl"
	}
	tpl.ExecuteTemplate(w, tn, nil)
}

// listHandler handles listing new pets for adoption.
func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Redirect(w, r, "/api/view/1234", http.StatusFound)
}

// staticd is the static asset and template service for Pet Shop.
func main() {
	var err error
	// Pre-load templates
	tpl, err = template.ParseFS(data.TemplateFS, "templates/*")
	if err != nil {
		log.Fatal(err)
	}
	// Configure http server
	hfs := http.FileServer(http.FS(data.StaticFS))
	http.Handle("/static/", hfs)
	http.HandleFunc("/api/list", listHandler)
	http.HandleFunc("/", templateHandler)
	// Start the server
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Println(err)
	}
}
