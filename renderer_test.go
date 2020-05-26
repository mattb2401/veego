package veego

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/square/go-jose.v2/json"
)

func TestRenderer_RespondWithJSON(t *testing.T) {
	rt := mux.NewRouter()
	router := NewRouter(rt)
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		renderer := NewRenderer(w, "json")
		renderer.JSON(map[string]interface{}{"code": 200}, 200)
	})
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("error occurred: %v", err.Error())
	}
	rec := httptest.NewRecorder()
	router.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected http code 200 but got %v", rec.Code)
	}
	type Resp struct {
		Code int `json:"code"`
	}
	var resp Resp
	bd, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("error decoding the body")
	}
	if err := json.Unmarshal(bd, &resp); err != nil {
		t.Errorf("error decoding the body to resp")
	}
	if resp.Code != 200 {
		t.Errorf("expected code 200 but got %v", resp.Code)
	}
}

func TestRenderer_RespondWithXML(t *testing.T) {
	rt := mux.NewRouter()
	router := NewRouter(rt)
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		renderer := NewRenderer(w, "xml")
		renderer.XML(map[string]interface{}{"code": 200}, "response", 200)
	})
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("error occurred: %v", err.Error())
	}
	rec := httptest.NewRecorder()
	router.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected http code 200 but got %v", rec.Code)
	}
	type Resp struct {
		XMLName xml.Name `xml:"response"`
		Code    int      `xml:"code"`
	}
	var resp Resp
	bd, err := ioutil.ReadAll(rec.Body)
	if err != nil {
		t.Errorf("error decoding the body")
	}
	if err := xml.Unmarshal(bd, &resp); err != nil {
		t.Errorf("error decoding the body to resp: %v", err.Error())
	}
	if resp.Code != 200 {
		t.Errorf("expected code 200 but got %v", resp.Code)
	}
}
