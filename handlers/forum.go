package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
)

func HandleForumCreate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		forum := &models.ForumSort{}
		err := unmarshalBody(r, forum)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createdForum, err := database.ForumCreate(env, forum)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(createdForum, w)
			return
		}
		switch err.(type) {
		case *models.ErrorForumAlreadyExists:
			err := err.(*models.ErrorForumAlreadyExists)
			w.WriteHeader(http.StatusConflict)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err.Forum, w)
		default:
			if !processErrorNotFound(w, err) {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	}
}

func HandleForumGet(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		forum, err := database.ForumGet(env, slug)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(forum, w)
			return
		}
		if !processErrorNotFound(w, err) {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandleForumUsers(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]

		limitParam := r.URL.Query().Get("limit")
		var limit uint64
		if limitParam != "" {
			var err error
			limit, err = strconv.ParseUint(limitParam, 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		since := r.URL.Query().Get("since")

		desc := r.URL.Query().Get("desc") == "true"

		users, err := database.ForumGetUsers(env, slug, since, limit, desc)

		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(users, w)
			return
		}
		if !processErrorNotFound(w, err) {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
