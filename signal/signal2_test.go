package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler1(t *testing.T) {
	w := httptest.NewRecorder() //fake http response writer
	r, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Handler() status = %d; wanted: %d", resp.StatusCode, 200) //do NOT continue the test
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; wanted: %q", contentType,
			"application/json") //DO CONTINUE with test, it is not so relevant to have json t this moment
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// fail test if not able to read body
		// will NOT put "wanted" as it is pretty obious we expect nil error
		t.Fatalf("iotil.ReadAll(resp.Bdy) err = %s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p) // unmarshall into a new Person struct
	if err != nil {
		// doesn't make sense to continue with property's tests if unable to unmarshall
		t.Fatalf("json.Unmarshall(resp.Body) err = %s", err)
	}

	if p.Age != 21 {
		// it is not so important to stop the rest of the tests
		t.Errorf("person.Age = %d; wanted: %d", p.Age, 21)
	}

	if p.Name != "Rafa" {
		t.Errorf("person.Name = %s; wanted: %s", p.Name, "Rafa")
	}
}
