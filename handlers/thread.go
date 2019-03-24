package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mailru/easyjson"

	"github.com/gorilla/mux"
)

func HandleThreadCreate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		forumSlug := vars["slug"]

		thread := &models.ThreadShort{}
		err := unmarshalBody(r, thread)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newThread, err := database.ThreadCreate(env, thread, forumSlug)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(newThread, w)
			return
		}
		switch err.(type) {
		case *models.ErrorThreadAlreadyExists:
			w.WriteHeader(http.StatusConflict)
			err := err.(*models.ErrorThreadAlreadyExists)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err.Thread, w)
		case *models.ErrorNotFound:
			w.WriteHeader(http.StatusNotFound)
			err := err.(*models.ErrorNotFound)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandleThreadList(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		forumSlug := vars["slug"]

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

		sinceParam := r.URL.Query().Get("since")
		var since time.Time
		if sinceParam != "" {
			var err error
			since, err = time.Parse("2006-01-02T15:04:05.000Z07:00", sinceParam)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		desc := r.URL.Query().Get("desc") == "true"

		threads, err := database.ThreadGetList(env, forumSlug, since, limit, desc)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(threads, w)
			return
		}
		switch err.(type) {
		case *models.ErrorNotFound:
			w.WriteHeader(http.StatusNotFound)
			err := err.(*models.ErrorNotFound)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandleThreadDetails(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, id := parseSlugOrId(mux.Vars(r)["slug_or_id"])
		thread, threadDBErr := database.ThreadGetBySlugOrId(env, slug, id)

		if threadDBErr == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(thread, w)
			w.WriteHeader(http.StatusOK)
			return
		}
		switch threadDBErr.(type) {
		case *models.ErrorNotFound:
			err := threadDBErr.(*models.ErrorNotFound)
			w.WriteHeader(http.StatusNotFound)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(`Bad request to thread details: %s`, threadDBErr.Error())
		}
	}
}
