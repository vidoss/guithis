package handlers

import (
	"appengine/datastore"
	"github.com/gorilla/mux"
	"github.com/vidoss/guithis/models"
	"net/http"
	"time"
	//"log"
)

func GetResource(c *AppContext, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		http.Error(w, "No ID!", http.StatusBadRequest)
	}

	GetResourceById(c, w, r, id)
}

func GetResourceById(c *AppContext, w http.ResponseWriter, r *http.Request, id string) {

	key, err := datastore.DecodeKey(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var resource models.Resource

	if err := datastore.Get(c.aeContext, key, &resource); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	resource.Id = id
	c.render.JSON(w, http.StatusOK, resource)
}

func GetAllResources(c *AppContext, w http.ResponseWriter, r *http.Request) {

	var resources []models.Resource

	q := datastore.NewQuery("Resource").Order("-Update").Limit(10)

	keys, err := q.GetAll(c.aeContext, &resources)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for i := 0; i < len(resources); i++ {
		resources[i].Id = keys[i].Encode()
	}

	c.render.JSON(w, http.StatusOK, resources)

}

func CreateResource(c *AppContext, w http.ResponseWriter, r *http.Request) {

	var resource models.Resource

	if readJson(r, &resource) {

		key := datastore.NewIncompleteKey(c.aeContext, "Resource", nil)

		resource.Update = time.Now()
		resource.Create = time.Now()

		key, err := datastore.Put(c.aeContext, key, &resource)
		resource.Id = key.Encode()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c.render.JSON(w, http.StatusOK, resource)
	}
}
