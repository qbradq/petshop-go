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

// templateHandler renders HTML templates
func templateHandler(response http.ResponseWriter, request *http.Request) {
	rn := path.Base(request.RequestURI)
	tn := "index.html.tpl"
	if rn != "/" && rn != "." {
		tn = rn + ".tpl"
	}
	tpl.ExecuteTemplate(response, tn, nil)
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
	http.HandleFunc("/", templateHandler)
	// Start the server
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Println(err)
	}
}
