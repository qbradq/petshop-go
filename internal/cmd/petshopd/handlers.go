package petshopd

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

// templateHandler renders HTML templates.
func templateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// TODO restrict template access to the front-facing ones
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
	// TODO input validation
	// File upload
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
	// Insert statement
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
	// Redirect to the view page
	http.Redirect(w, r, fmt.Sprintf("/view.html?p=%d", id), http.StatusFound)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		name        string
		description string
	)
	v := r.URL.Query()
	va := v["p"]
	if len(va) < 1 {
		http.Error(w, "no pet specified", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(va[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = statements["view"].QueryRow(id).Scan(&name, &description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pet := &Pet{
		ID:          id,
		Name:        name,
		Description: description,
	}
	tpl.ExecuteTemplate(w, "view.html.tpl", pet)
}
