package handlers

import (
	"appengine/datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vidoss/guithis/context"
	"github.com/vidoss/guithis/models"
	"net/http"
	"time"
	//"log"
)

func GetResource(c *context.AppContext, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		c.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "No ID!"})
		return
	}

	GetResourceById(c, w, r, id)
}

func GetResourceById(c *context.AppContext, w http.ResponseWriter, r *http.Request, id string) {

	key, err := datastore.DecodeKey(id)
	if err != nil {
		c.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var resource models.Resource

	if err := datastore.Get(c.GaeContext, key, &resource); err != nil {
		c.Render.JSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	resource.Id = id
	c.Render.JSON(w, http.StatusOK, resource)
}

func GetAllResources(c *context.AppContext, w http.ResponseWriter, r *http.Request) {

	var resources []models.Resource

	q := datastore.NewQuery("Resource").Order("-Update").Limit(10)

	keys, err := q.GetAll(c.GaeContext, &resources)

	if err != nil {
		c.Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	for i := 0; i < len(resources); i++ {
		resources[i].Id = keys[i].Encode()
	}

	c.Render.JSON(w, http.StatusOK, resources)

}

func CreateResource(c *context.AppContext, w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var resource models.Resource

	if err := decoder.Decode(&resource); err != nil {
		c.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	key := datastore.NewIncompleteKey(c.GaeContext, "Resource", nil)

	resource.Update = time.Now()
	resource.Create = time.Now()

	key, err := datastore.Put(c.GaeContext, key, &resource)
	resource.Id = key.Encode()

	if err != nil {
		c.Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.Render.JSON(w, http.StatusOK, resource)
}
