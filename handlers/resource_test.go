package handlers

import (
	"appengine/aetest"
	"encoding/json"
	"errors"
	"github.com/unrolled/render"
	"github.com/vidoss/guithis/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResource(t *testing.T) {

	aeContext, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer aeContext.Close()

	context := &AppContext{
		aeContext: aeContext,
		render:    render.New(render.Options{}),
	}

	if key, err := testCreateResource(context, t); err == nil {
		testGetResource(context, t, key)
		testGetAllResources(context, t, key)
	}
}

func testCreateResource(c *AppContext, t *testing.T) (key string, err error) {

	data := `{"name":"Customer","description":"Customer Description"}`

	r, _ := http.NewRequest("POST", "/resource", strings.NewReader(data))
	w := httptest.NewRecorder()

	CreateResource(c, w, r)

	if w.Code != http.StatusOK {
		t.Errorf("POST /resource Failed, Expected 200, Got %v", w.Code)
		return "", errors.New("Create Failed")
	}

	var resource models.Resource
	if err := readResponse(w, &resource); err != nil {
		return "", err
	}

	return resource.Id, nil
}

func testGetResource(c *AppContext, t *testing.T, k string) {

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

	if resource.Id != k {
		t.Errorf("Resource Id does not match: Expected %v, Got %v\n", k, resource.Id)
	}
}

func testGetAllResources(c *AppContext, t *testing.T, k string) {

	r, _ := http.NewRequest("GET", "/resource", nil)
	w := httptest.NewRecorder()

	GetAllResources(c, w, r)

	if w.Code != http.StatusOK {
		t.Error("GET /resource Failed, status not ok...")
	}

	var resArr []models.Resource
	if err := readResponse(w, &resArr); err != nil {
		t.Error(err)
	}

	if len(resArr) != 1 {
		t.Error("Resource not created...")
		return
	}

	if resArr[0].Name != "Customer" {
		t.Errorf("Expected: Customer, Got %v\n", resArr[0].Name)
	}

	if resArr[0].Id != k {
		t.Errorf("Expected: %v, Got %v\n", resArr[0].Id)
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
