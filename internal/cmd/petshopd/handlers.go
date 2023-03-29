package petshopd

import (
	"fmt"
	"io"
	"net/http"
	"path"
)

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
	if err := r.ParseMultipartForm(1024 * 1024 * 10); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name := r.PostForm.Get("name")
	description := r.PostForm.Get("description")
	pf, _, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer pf.Close()
	pd, err := io.ReadAll(pf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := statements["list"].Exec(name, description, pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/api/view/%d", id), http.StatusFound)
}
