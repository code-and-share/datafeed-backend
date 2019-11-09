package main

// Based on https://github.com/le4ndro/gowt

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

// Object struct
type Object struct {
	Id      int
	Name    string
	Content string
}

// JSON definition of Object inside Phase
type PhaseObjects struct {
	Object   string `json:"object"`
	Position string `json:"position"`
}

// Phase struct
type Phase struct {
	Id      int
	Name    string
	Objects []PhaseObjects
}

// Path struct
type Path struct {
	Id         int
	Name       string
	PhaseOrder int
	PhaseId    int
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbServer := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbServer+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("templates/*"))

//Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	var aux string
	tmpl.ExecuteTemplate(w, "Index", aux)
}

//Objects Main handler
func Objects(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM objects ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	object := Object{}
	res := []Object{}

	for selDB.Next() {
		var id int
		var name, content string
		err := selDB.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | name " + name + " | content " + content)

		object.Id = id
		object.Name = name
		object.Content = content
		res = append(res, object)
	}
	tmpl.ExecuteTemplate(w, "Objects", res)
	defer db.Close()
}

//Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM objects WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	object := Object{}

	for selDB.Next() {
		var id int
		var name, content string
		err := selDB.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | name " + name + " | content " + content)

		object.Id = id
		object.Name = name
		object.Content = content
	}
	tmpl.ExecuteTemplate(w, "Show", object)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM objects WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	object := Object{}

	for selDB.Next() {
		var id int
		var name, content string
		err := selDB.Scan(&id, &name, &content)
		if err != nil {
			panic(err.Error())
		}

		object.Id = id
		object.Name = name
		object.Content = content
	}

	tmpl.ExecuteTemplate(w, "Edit", object)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		content := r.FormValue("content")
		insForm, err := db.Prepare("INSERT INTO objects (name, content) VALUES (?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, content)
		log.Println("Insert Data: name " + name + " | content " + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		content := r.FormValue("content")
		insForm, err := db.Prepare("UPDATE objects SET name=?, content=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, content)
		log.Println("UPDATE Data: name " + name + " | content " + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	object := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM objects WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(object)
	log.Println("DELETE " + object)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/objects", Objects)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
