package handlers

import (
	"appengine"
	"appengine/aetest"
	"encoding/json"
	"github.com/vidoss/guithis/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"errors"
)

func TestResource(t *testing.T) {

	c, err := aetest.NewContext(nil)

	if err != nil {
		t.Fatal(err)
	}

	defer c.Close()

	if key, err := testCreateResource(c, t); err == nil {
		testGetResource(c, t, key)
		testGetAllResources(c, t)
	}
}

func testCreateResource(c appengine.Context, t *testing.T) (key string, err error) {

	data := `{"name":"Customer","description":"Customer Description"}`

	r, _ := http.NewRequest("POST", "/resource", strings.NewReader(data))
	w := httptest.NewRecorder()

	CreateResource(c, w, r)

	if w.Code != http.StatusOK {
		t.Errorf("POST /resource Failed, Expected 200, Got %v",w.Code)
		return "",  errors.New("Create Failed")
	}

	var resource models.Resource
	if err := readResponse(w, &resource); err != nil {
		return "", err
	}

	return resource.Id, nil
}

func testGetResource(c appengine.Context, t *testing.T, k string) {

	r, _ := http.NewRequest("POST", "/resource/"+k, nil)
	w := httptest.NewRecorder()

	GetResourceById(c, w, r, k)

	if w.Code != http.StatusOK {
		t.Errorf("GET /resource/%v Failed, Expected 200, Got %v", k, w.Code)
		return
	}

	var resource models.Resource
	if err := readResponse(w, &resource); err != nil {
		t.Error(err)
	}
}

func testGetAllResources(c appengine.Context, t *testing.T) {

	r, _ := http.NewRequest("GET", "/resource", nil)
	w := httptest.NewRecorder()

	GetAllResources(c, w, r)

	if w.Code != http.StatusOK {
		t.Error("GET /resource Failed, status not ok...")
	}

	var resourcesArr []models.Resource;
	if err := readResponse(w, &resourcesArr); err != nil {
		t.Error(err)
	}

	if len(resourcesArr) == 0 {
		t.Error("Resource not created...")
		return
	}
	if resourcesArr[0].Name != "Customer" {
		t.Error("Not the created resource...")
	}

}

func readResponse(w *httptest.ResponseRecorder, v interface{}) error {

	body, err := ioutil.ReadAll(w.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}
