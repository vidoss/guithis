package app

import (
	"appengine"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	h "github.com/vidoss/guithis/handlers"
	"log"
	"net/http"
)

type appHandler struct {
	c  *h.AppContext
	fn func(*h.AppContext, http.ResponseWriter, *http.Request)
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ah.c.SetAppEngineContext(appengine.NewContext(r))

	log.Println("%v", ah.c)
	ah.fn(ah.c, w, r)
}

func init() {

	// Init context
	context := h.NewAppContext()

	log.Println("Setting up handlers...")

	r := mux.NewRouter()
	n := negroni.Classic()

	r.Handle("/resource", appHandler{context, h.GetAllResources}).Methods("GET")
	r.Handle("/resource/{id}", appHandler{context, h.GetResource}).Methods("GET")
	r.Handle("/resource", appHandler{context, h.CreateResource}).Methods("POST")

	n.UseHandler(r)

	http.Handle("/", n)

}
