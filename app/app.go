package app

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	h "github.com/vidoss/guithis/handlers"
	"log"
	"net/http"
)

func init() {

	log.Println("Setting up handlers...")

	r := mux.NewRouter()
	n := negroni.Classic()

	r.Handle("/resource", h.ContextHandler{h.GetAllResources}).Methods("GET")
	r.Handle("/resource/{id}", h.ContextHandler{h.GetResource}).Methods("GET")
	r.Handle("/resource", h.ContextHandler{h.CreateResource}).Methods("POST")

	n.UseHandler(r)

	http.Handle("/", n)

}
