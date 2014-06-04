package handlers

import (
	"github.com/vidoss/guithis/models"
	"net/http"
)

var resources []models.Resource

func GetAllResources(w http.ResponseWriter, r *http.Request) {
	writeJson(w, resources)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource = models.Resource{}

	if readJson(r, &resource) {
		resources = append(resources, resource)
		writeJson(w, resource)
	}
}
