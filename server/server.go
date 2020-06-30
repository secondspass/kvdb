package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/secondspass/kvdb/raft"
)

// kvstore is a var of the interface. TODO: maybe change this to a raft.DB interface or
// something?
var kvstore raft.DB

func get(w http.ResponseWriter, r *http.Request) {
	var data string
	key := r.FormValue("key")
	data = kvstore.Get(key)
	if data == "" {
		log.Printf("no data for key: %s", key)
		data = "no data for key"
	}
	tmpl, err := template.ParseFiles("server/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, data)
}

func put(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("val")
	err := kvstore.Put(key, value)
	if err != nil {
		log.Print(err)
	}
	tmpl, err := template.ParseFiles("server/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("server/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

// Start starts up the API server to interact with the db
func Start() {
	// TODO: change this db.Open to a raft middleware implementation of the db, that
	// will also return a value that implements the db.DB interface. But calling get
	// and put on it will do the raft consensus stuff and all that instead of directly
	// storing.
	temp, err := raft.Open("my.db")
	if err != nil {
		log.Fatal(err)
	}
	kvstore = temp

	http.HandleFunc("/get", get)
	http.HandleFunc("/put", put)
	http.HandleFunc("/", home)
	fmt.Println("Serving on localhost:8090")
	http.ListenAndServe(":8090", nil)
}
