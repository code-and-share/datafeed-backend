package main

// Based on https://github.com/le4ndro/gowt

import (
	"database/sql"
	"encoding/json"
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

//-----------------------------------------------------------
// Functions to handle Objects
//-----------------------------------------------------------
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

func ObjectsShow(w http.ResponseWriter, r *http.Request) {
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
	tmpl.ExecuteTemplate(w, "Objects_Show", object)
	defer db.Close()
}

func ObjectsNew(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Objects_New", nil)
}

func ObjectsEdit(w http.ResponseWriter, r *http.Request) {
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

	tmpl.ExecuteTemplate(w, "Objects_Edit", object)
	defer db.Close()
}

func ObjectsInsert(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/objects", 301)
}

func ObjectsUpdate(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/objects", 301)
}

func ObjectsDelete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	object := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM objects WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(object)
	log.Println("DELETE " + object)
	defer db.Close()
	http.Redirect(w, r, "/objects", 301)
}

//-----------------------------------------------------------
// Functions to handle Phases
//-----------------------------------------------------------
func Phases(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM phases ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	phase := Phase{}
	res := []Phase{}

	for selDB.Next() {
		var id int
		var name, objects_json string
		err := selDB.Scan(&id, &name, &objects_json)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | name " + name + " | objects " + objects_json)

		objects := make([]PhaseObjects, 0)

		err = json.Unmarshal([]byte(objects_json), &objects)
		if err != nil {
			log.Println("Error decoding JSON: " + err.Error())
		}
		phase.Id = id
		phase.Name = name
		phase.Objects = objects
		res = append(res, phase)
	}
	tmpl.ExecuteTemplate(w, "Phases", res)
	defer db.Close()
}

func PhasesShow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phases WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	phase := Phase{}

	for selDB.Next() {
		var id int
		var name, objects_json string
		err := selDB.Scan(&id, &name, &objects_json)
		if err != nil {
			panic(err.Error())
		}
		objects := make([]PhaseObjects, 0)
		err = json.Unmarshal([]byte(objects_json), &objects)
		if err != nil {
			log.Println("Error decoding JSON: " + err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | name " + name + " | objects " + objects_json)

		phase.Id = id
		phase.Name = name
		phase.Objects = objects
	}
	tmpl.ExecuteTemplate(w, "Phases_Show", phase)
	defer db.Close()
}

func PhasesNew(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Phases_New", nil)
}

func PhasesEdit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phases WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	phase := Phase{}

	for selDB.Next() {
		var id int
		var name, objects_json string
		err := selDB.Scan(&id, &name, &objects_json)
		if err != nil {
			panic(err.Error())
		}

		objects := make([]PhaseObjects, 0)
		err = json.Unmarshal([]byte(objects_json), &objects)
		if err != nil {
			log.Println("Error decoding JSON: " + err.Error())
		}

		phase.Id = id
		phase.Name = name
		phase.Objects = objects
	}

	tmpl.ExecuteTemplate(w, "Phases_Edit", phase)
	defer db.Close()
}

//-----------------------------------------------------------
// Functions to handle Paths
//-----------------------------------------------------------
func Paths(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM paths ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	path := Path{}
	res := []Path{}

	for selDB.Next() {
		var id, phase_order, phase_id int
		var name string
		err := selDB.Scan(&id, &name, &phase_order, &phase_id)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Listing Row: Id " + string(id) + " | name " + name + " | phase_order " + string(phase_order) + " | phase_id " + string(phase_id))

		path.Id = id
		path.Name = name
		path.PhaseOrder = phase_order
		path.PhaseId = phase_id
		res = append(res, path)
	}
	tmpl.ExecuteTemplate(w, "Paths", res)
	defer db.Close()
}
func PathsShow(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM paths WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	path := Path{}

	for selDB.Next() {
		var id, phase_order, phase_id int
		var name string
		err := selDB.Scan(&id, &name, &phase_order, &phase_id)
		if err != nil {
			panic(err.Error())
		}

		log.Println("Showing Row: Id " + string(id) + " | name " + name + " | phase_order " + string(phase_order) + " | phase_id " + string(phase_id))

		path.Id = id
		path.Name = name
		path.PhaseOrder = phase_order
		path.PhaseId = phase_id
	}
	tmpl.ExecuteTemplate(w, "Paths_Show", path)
	defer db.Close()
}

func PathsNew(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Paths_New", nil)
}

func PathsEdit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM paths WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}

	path := Path{}

	for selDB.Next() {
		var id, phase_order, phase_id int
		var name string
		err := selDB.Scan(&id, &name, &phase_order, &phase_id)
		if err != nil {
			panic(err.Error())
		}

		path.Id = id
		path.Name = name
		path.PhaseOrder = phase_order
		path.PhaseId = phase_id
	}

	tmpl.ExecuteTemplate(w, "Paths_Edit", path)
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/objects", Objects)
	http.HandleFunc("/objects_show", ObjectsShow)
	http.HandleFunc("/objects_new", ObjectsNew)
	http.HandleFunc("/objects_edit", ObjectsEdit)
	http.HandleFunc("/objects_insert", ObjectsInsert)
	http.HandleFunc("/objects_update", ObjectsUpdate)
	http.HandleFunc("/objects_delete", ObjectsDelete)
	http.HandleFunc("/phases", Phases)
	http.HandleFunc("/phases_show", PhasesShow)
	http.HandleFunc("/phases_new", PhasesNew)
	http.HandleFunc("/paths", Paths)
	http.HandleFunc("/paths_show", PathsShow)
	http.HandleFunc("/paths_new", PathsNew)
	http.ListenAndServe(":8080", nil)
}
