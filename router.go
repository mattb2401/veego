package veego

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Route *mux.Router
}

func NewRouter(router *mux.Router) *Router {
	return &Router{Route: router}
}

func (r *Router) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Route.HandleFunc(path, handler).Methods("POST")
}

func (r *Router) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Route.HandleFunc(path, handler).Methods("GET")
}
