package router

import (
	"ForumDB/handlers"
	"ForumDB/models"
	"net/http"
)

type Route struct {
	name       string
	path       string
	method     string
	handleFunc func(env *models.Env) http.HandlerFunc
}

var routes = []Route{
	{
		"UserCreate",
		"/user/{nickname}/create",
		"POST",
		handlers.HandleUserCreate,
	},
	{
		"UserGet",
		"/user/{nickname}/profile",
		"GET",
		handlers.HandleUserGet,
	},
	{
		"UserUpdate",
		"/user/{nickname}/profile",
		"POST",
		handlers.HandleUserUpdate,
	},
}
