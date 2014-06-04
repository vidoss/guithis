package app

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vidoss/guithis/handlers"
	"log"
	"net/http"
)

func init() {

	log.Println("Setting up handlers...")

	r := mux.NewRouter()
	n := negroni.Classic()

	r.HandleFunc("/resource", handlers.GetAllResources).Methods("GET")
	r.HandleFunc("/resource", handlers.CreateResource).Methods("POST")

	n.UseHandler(r)

	http.Handle("/", n)

}
