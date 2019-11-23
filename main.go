package main

// Based on https://github.com/le4ndro/gowt

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/lib/pq"
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

func dbConnPostgres() (db *sql.DB) {
	dbDriver := "postgres"
	dbUser := os.Getenv("DB_PSQL_USER")
	dbPass := os.Getenv("DB_PSQL_PASS")
	dbName := os.Getenv("DB_PSQL_NAME")
	dbHost := os.Getenv("DB_PSQL_HOST")
	dbPort := os.Getenv("DB_PSQL_PORT")
	db, err := sql.Open(dbDriver, "postgres://"+dbUser+":"+dbPass+"@"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable")
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
	db := dbConnPostgres()
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
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM objects WHERE id=$1", nId)
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
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM objects WHERE id=$1", nId)
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
	db := dbConnPostgres()
	if r.Method == "POST" {
		name := r.FormValue("name")
		content := name + ".png"
		// Get file
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("content_file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		//// write this byte array to our temporary file
		//tempFile.Write(fileBytes)
		err = ioutil.WriteFile("./files/"+content, fileBytes, 0644)
		if err != nil {
			fmt.Println(err)
		}

		// Add entry to our DB
		insForm, err := db.Prepare("INSERT INTO objects (name, content) VALUES ($1, $2)")
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
	db := dbConnPostgres()
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		content := name + ".png"
		// Get file
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("content_file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		//// write this byte array to our temporary file
		//tempFile.Write(fileBytes)
		err = ioutil.WriteFile("./files/"+content, fileBytes, 0644)
		if err != nil {
			fmt.Println(err)
		}

		insForm, err := db.Prepare("UPDATE objects SET name=$1, content=$2 WHERE id=$3")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, content, id)
		log.Println("UPDATE Data: name " + name + " | content " + content)
	}
	defer db.Close()
	http.Redirect(w, r, "/objects", 301)
}

func ObjectsDelete(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	object := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM objects WHERE id=$1")
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
	db := dbConnPostgres()
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
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phases WHERE id=$1", nId)
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
	db := dbConnPostgres()
	selDB, err := db.Query("SELECT name FROM objects ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	var object Object
	res := []Object{}

	for selDB.Next() {
		var name string
		err := selDB.Scan(&name)
		if err != nil {
			panic(err.Error())
		}

		object.Name = name
		res = append(res, object)
	}
	defer db.Close()
	tmpl.ExecuteTemplate(w, "Phases_New", res)
}

func PhasesEdit(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM phases WHERE id=$1", nId)
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

func PhasesInsert(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	if r.Method == "POST" {
		name := r.FormValue("name")
		objects := r.FormValue("objects")
		insForm, err := db.Prepare("INSERT INTO phases (name, objects) VALUES ($1, $2)")
		if err != nil {
			panic(err.Error())
		}
		_, err = insForm.Exec(name, objects)
		if err != nil {
			panic(err.Error())
		}
		log.Println("Insert Data: name " + name + " | objects " + objects)
	}
	defer db.Close()
	http.Redirect(w, r, "/phases", 301)
}

func PhasesUpdate(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		objects := r.FormValue("objects")
		insForm, err := db.Prepare("UPDATE phases SET name=$1, objects=$2 WHERE id=$3")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, objects, id)
		log.Println("UPDATE Data: name " + name + " | objects " + objects)
	}
	defer db.Close()
	http.Redirect(w, r, "/phases", 301)
}

func PhasesDelete(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	phase := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM phases WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(phase)
	log.Println("DELETE " + phase)
	defer db.Close()
	http.Redirect(w, r, "/phases", 301)
}

//-----------------------------------------------------------
// Functions to handle Paths
//-----------------------------------------------------------
func Paths(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
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
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM paths WHERE id=$1", nId)
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
	db := dbConnPostgres()
	selDB, err := db.Query("SELECT id, name FROM phases ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	phase := Phase{}
	res := []Phase{}

	for selDB.Next() {
		var id int
		var name string
		err := selDB.Scan(&id, &name)
		if err != nil {
			panic(err.Error())
		}

		phase.Id = id
		phase.Name = name
		res = append(res, phase)
	}
	tmpl.ExecuteTemplate(w, "Paths_New", res)
	defer db.Close()
}

func PathsEdit(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM paths WHERE id=$1", nId)
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

func PathsInsert(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	if r.Method == "POST" {
		name := r.FormValue("name")
		phase_order := r.FormValue("phase_order")
		phase_id, _ := strconv.Atoi(strings.Split(r.FormValue("phase_id"), `->`)[0])
		insForm, err := db.Prepare("INSERT INTO paths (name, phase_order, phase_id) VALUES ($1, $2, $3)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, phase_order, phase_id)
		log.Println("Insert Data: name " + name + " | phase_order " + phase_order + " | phase_id " + strconv.Itoa(phase_id))
	}
	defer db.Close()
	http.Redirect(w, r, "/paths", 301)
}

func PathsUpdate(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		phase_order := r.FormValue("phase_order")
		phase_id := r.FormValue("phase_id")
		insForm, err := db.Prepare("UPDATE paths SET name=$1, phase_order=$2, phase_id=$3 WHERE id=$4")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, phase_order, phase_id, id)
		log.Println("UPDATE Data: name " + name + " | phase_order " + phase_order + " | phase_id " + phase_id)
	}
	defer db.Close()
	http.Redirect(w, r, "/paths", 301)
}

func PathsDelete(w http.ResponseWriter, r *http.Request) {
	db := dbConnPostgres()
	path := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM paths WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(path)
	log.Println("DELETE " + path)
	defer db.Close()
	http.Redirect(w, r, "/paths", 301)
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
	http.HandleFunc("/phases_edit", PhasesEdit)
	http.HandleFunc("/phases_insert", PhasesInsert)
	http.HandleFunc("/phases_update", PhasesUpdate)
	http.HandleFunc("/phases_delete", PhasesDelete)
	http.HandleFunc("/paths", Paths)
	http.HandleFunc("/paths_show", PathsShow)
	http.HandleFunc("/paths_new", PathsNew)
	http.HandleFunc("/paths_edit", PathsEdit)
	http.HandleFunc("/paths_insert", PathsInsert)
	http.HandleFunc("/paths_update", PathsUpdate)
	http.HandleFunc("/paths_delete", PathsDelete)
	//http.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files/"))))
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./files/"))))
	http.ListenAndServe(":8080", nil)
}
