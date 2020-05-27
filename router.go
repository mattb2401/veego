package veego

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter(r *mux.Router) *Router {
	return &Router{Router: r}
}
//x
func (r *Router) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("POST")
}

func (r *Router) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("GET")
}

func (r *Router) Put(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("PUT")
}

func (r *Router) Delete(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("DELETE")
}