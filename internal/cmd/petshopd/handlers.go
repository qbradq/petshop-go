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
	pf, handle, err := r.FormFile("picture")
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
	// File attributes
	mime := handle.Header.Get("Content-Type")
	ext := path.Ext(handle.Filename)
	// Insert statement
	res, err := statements["list"].Exec(name, description, mime, ext, pd)
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

// viewHandler handles viewing a pet for adoption.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var (
		name        string
		description string
		ext         string
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
	err = statements["view"].QueryRow(id).Scan(&name, &description, &ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pet := &Pet{
		ID:               id,
		Name:             name,
		Description:      description,
		PictureExtension: ext,
	}
	if err := tpl.ExecuteTemplate(w, "view.html.tpl", pet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// imageHandler serves pet images.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	var pd = make([]byte, 0)
	var mime string
	ids := r.URL.Query().Get("p")
	id, err := strconv.Atoi(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = statements["image"].QueryRow(id).Scan(&mime, &pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", mime)
	w.WriteHeader(http.StatusOK)
	w.Write(pd)
}
