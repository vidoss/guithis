package app

import (
	"appengine"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/vidoss/guithis/context"
	h "github.com/vidoss/guithis/handlers"
	"github.com/vidoss/guithis/views"
	"log"
	"net/http"
)

type appHandler struct {
	c  *context.AppContext
	fn func(*context.AppContext, http.ResponseWriter, *http.Request)
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ah.c.GaeContext = appengine.NewContext(r)

	ah.fn(ah.c, w, r)
}

func init() {

	// Init context
	context := context.NewAppContext(render.Options{
		Directory: "../templates",
		Layout:    "layout",
	})

	log.Println("Setting up handlers...")

	r := mux.NewRouter()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("../public")))

	r.Handle("/", appHandler{context, views.HomeHandler}).Methods("GET")

	r.Handle("/resource", appHandler{context, h.GetAllResources}).Methods("GET")
	r.Handle("/resource/{id}", appHandler{context, h.GetResource}).Methods("GET")
	r.Handle("/resource", appHandler{context, h.CreateResource}).Methods("POST")

	n.UseHandler(r)

	http.Handle("/", n)

}
