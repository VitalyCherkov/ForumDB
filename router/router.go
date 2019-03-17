package router

import (
	"ForumDB/models"
	"fmt"

	"github.com/gorilla/mux"
)

func Init(env *models.Env) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	for _, route := range routes {
		fmt.Printf("go:init %v\t%v\t%v\n", route.name, route.method, route.path)
		handler := route.handleFunc(env)

		api.HandleFunc(route.path, handler).Methods(route.method)
	}

	return r
}
