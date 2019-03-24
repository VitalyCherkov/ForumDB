package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"

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
		case *models.ErrorNotFound:
			processErrorNotFound(w, err)
		default:
			fmt.Println("forum create bad request: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
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
		switch err.(type) {
		case *models.ErrorNotFound:
			processErrorNotFound(w, err)
		default:
			fmt.Println("Forum bad request: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
