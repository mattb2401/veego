package veego

import (
	"net/http"

	"github.com/gorilla/mux"
)

type router struct {
	Router *mux.Router
}

func NewRouter(r *mux.Router) *router {
	return &router{Router: r}
}
//x
func (r *router) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("POST")
}

func (r *router) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("GET")
}

func (r *router) Put(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("PUT")
}

func (r *router) Delete(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("DELETE")
}