package handlers

import (
	"appengine/aetest"
	"encoding/json"
	"github.com/vidoss/guithis/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateResource(t *testing.T) {

	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	data := `{"name":"Customer","description":"Customer Description"}`

	r, _ := http.NewRequest("POST", "/resource", strings.NewReader(data))
	w := httptest.NewRecorder()

	CreateResource(w, r)

	if w.Code != http.StatusOK {
		t.Error("POST /resource Failed, status not ok...")
	}

}

func TestGetAllResources(t *testing.T) {
	r, _ := http.NewRequest("GET", "/resource", nil)
	w := httptest.NewRecorder()

	GetAllResources(w, r)

	if w.Code != http.StatusOK {
		t.Error("GET /resource Failed, status not ok...")
	}

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal("Unable to read /resource")
	}
	var resourcesArr = make([]models.Resource, 1)
	json.Unmarshal(body, &resourcesArr)
	if len(resourcesArr) == 0 {
		t.Error("Resource not created...")
	}
	if resourcesArr[0].Name != "Customer" {
		t.Error("Not the created resource...")
	}
}
