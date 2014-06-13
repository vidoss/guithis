package handlers

import (
	"appengine"
	"appengine/datastore"
	"github.com/vidoss/guithis/models"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	//"log"
)

func GetResource(c appengine.Context, w http.ResponseWriter, r *http.Request) {

		  vars := mux.Vars(r)
		  id, ok := vars["id"]

		  if !ok {
				http.Error(w, "No ID!", http.StatusBadRequest)
		  }

		  GetResourceById(c, w, r, id)
}

func GetResourceById(c appengine.Context, w http.ResponseWriter, r *http.Request, id string) {

		  key, err := datastore.DecodeKey(id)
		  if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
		  }

		  var resource models.Resource

		  if err := datastore.Get(c, key, &resource); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
		  }

		  resource.Id = id
		  writeJson(w, resource)
}

func GetAllResources(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	var resources []models.Resource ;

	q := datastore.NewQuery("Resource").Order("-Update").Limit(10)

	keys, err := q.GetAll(c, &resources)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for i := 0; i < len(resources); i++ {
		resources[i].Id = keys[i].Encode()
	}

	writeJson(w, resources)

}

func CreateResource(c appengine.Context, w http.ResponseWriter, r *http.Request) {

	var resource models.Resource

	if readJson(r, &resource) {

		key := datastore.NewIncompleteKey(c, "Resource", nil)

		resource.Update = time.Now()
		resource.Create = time.Now()

		key, err := datastore.Put(c, key, &resource)
		resource.Id = key.Encode()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJson(w, resource)
	}
}

