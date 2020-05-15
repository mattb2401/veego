package veego

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BaseRouter struct {
	Router *mux.Router
}

func NewRouter(router *mux.Router) *BaseRouter {
	return &BaseRouter{Router: router}
}

func (r *BaseRouter) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("POST")
}

func (r *BaseRouter) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.Router.HandleFunc(path, handler).Methods("GET")
}
