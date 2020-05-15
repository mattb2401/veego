package veego

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRouter_Get(t *testing.T) {
	rt := mux.NewRouter()
	router := NewRouter(rt)
	router.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		return
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
}

func TestRouter_Post(t *testing.T) {
	rt := mux.NewRouter()
	router := NewRouter(rt)
	router.Post("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		return
	})
	req, err := http.NewRequest("POST", "/test", nil)
	if err != nil {
		t.Errorf("error occurred: %v", err.Error())
	}
	rec := httptest.NewRecorder()
	router.Router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected http code 200 but got %v", rec.Code)
	}
}
