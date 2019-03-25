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

	{
		"ForumCreate",
		"/forum/create",
		"POST",
		handlers.HandleForumCreate,
	},
	{
		"ForumGet",
		"/forum/{slug}/details",
		"GET",
		handlers.HandleForumGet,
	},
	{
		"ForumUsers",
		"/forum/{slug}/users",
		"GET",
		handlers.HandleForumUsers,
	},

	{
		"ThreadCreate",
		"/forum/{slug}/create",
		"POST",
		handlers.HandleThreadCreate,
	},
	{
		"ThreadList",
		"/forum/{slug}/threads",
		"GET",
		handlers.HandleThreadList,
	},
	{
		"ThreadCreate",
		"/thread/{slug_or_id}/details",
		"GET",
		handlers.HandleThreadDetails,
	},
	{
		"ThreadDetails",
		"/thread/{slug_or_id}/details",
		"POST",
		handlers.HandleThreadUpdate,
	},
	{
		"ThreadDoVote",
		"/thread/{slug_or_id}/vote",
		"POST",
		handlers.HandleThreadDoVote,
	},

	{
		"PostListCreate",
		"/thread/{slug_or_id}/create",
		"POST",
		handlers.HandlePostListCreate,
	},
	{
		"PostDetail",
		"/post/{id}/details",
		"GET",
		handlers.HandlePostDetail,
	},
	{
		"PostDetail",
		"/post/{id}/details",
		"POST",
		handlers.HandlePostUpdate,
	},

	{
		"ServiceStatus",
		"/service/status",
		"GET",
		handlers.HandleServiceStatus,
	},
	{
		"ServiceClear",
		"/service/clear",
		"POST",
		handlers.HandleServiceClear,
	},
}
