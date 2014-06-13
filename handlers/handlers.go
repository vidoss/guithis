package handlers

import (
	"appengine"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

/* https://github.com/DenverGophers/talks/blob/master/2013-04/mgo/example_6/read_json.go */

// readJson will parses the JSON-encoded data in the http request and store the result in v
func readJson(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	var (
		body []byte
		err  error
	)

	body, err = ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request body %v", err)
		return false
	}

	return true
}

func writeJson(w http.ResponseWriter, v interface{}) {

	if data, err := json.Marshal(v); err != nil {
		log.Printf("Error marshalling json: %v", err)
	} else {
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

/* http://stackoverflow.com/questions/19407343/how-can-i-unit-test-google-app-engine-go-http-handlers/23121756#23121756 */

type ContextHandler struct {
	Real func(appengine.Context, http.ResponseWriter, *http.Request)
}

func (f ContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	f.Real(c, w, r)
}
