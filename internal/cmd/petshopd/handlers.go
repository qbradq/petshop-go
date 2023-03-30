package petshopd

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	ext := path.Ext(handle.Filename)
	pd, err := io.ReadAll(pf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Insert statement
	res, err := statements["list"].Exec(name, description, ext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Save the image to the file system
	err = os.WriteFile(path.Join("image", fmt.Sprintf("%d%s", id, ext)), pd, 0777)
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
		http.Error(w, "no pet specified", http.StatusBadRequest)
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
