package petshopd

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qbradq/petshop-go/data"
)

// tpl is the pre-parsed template file system
var tpl *template.Template

// db is the global database connector, is concurrent-safe and manages an
// internal connection pool
var db *sql.DB

// statements is the global cache of prepared statements
var statements = map[string]*sql.Stmt{}

// prep prepares a statement
func prep(name, statement string) {
	s, err := db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	statements[name] = s
}

// Main is the entry point for the Pet Shop service.
func Main() {
	var err error
	// Pre-load templates
	tpl, err = template.ParseFS(data.TemplateFS, "templates/*")
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the database
	db, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS pets (" +
			"uid INTEGER PRIMARY KEY AUTOINCREMENT," +
			"name VARCHAR(64) NULL," +
			"description VARCHAR(1024) NULL," +
			"picture_mime VARCHAR(64) NULL," +
			"picture_ext VARCHAR(32) NULL," +
			"picture BLOB NULL" +
			");")
	if err != nil {
		log.Fatal(err)
	}
	// Prepare statements
	prep("list", "INSERT INTO pets (name, description, picture_mime, picture_ext, picture) VALUES (?,?,?,?,?)")
	prep("view", "SELECT name, description, picture_ext FROM pets WHERE uid=?")
	prep("image", "SELECT picture_mime, picture FROM pets WHERE uid=?")
	// Configure http server
	hfs := http.FileServer(http.FS(data.StaticFS))
	http.Handle("/static/", hfs)
	http.HandleFunc("/api/test/*", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
	})
	http.HandleFunc("/api/image", imageHandler)
	http.HandleFunc("/api/list", listHandler)
	http.HandleFunc("/view.html", viewHandler)
	http.HandleFunc("/", templateHandler)
	// Start the server
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Println(err)
	}
}
