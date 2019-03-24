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
			processErrorNotFound(w, err)
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
			processErrorNotFound(w, err)
		default:
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandleThreadDetails(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, id := parseSlugOrId(r)
		thread, threadDBErr := database.ThreadGetBySlugOrId(env, slug, id)

		if threadDBErr == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(thread, w)
			return
		}
		switch threadDBErr.(type) {
		case *models.ErrorNotFound:
			processErrorNotFound(w, threadDBErr)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(`Bad request to thread details: %s`, threadDBErr.Error())
		}
	}
}

func HandleThreadDoVote(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, id := parseSlugOrId(r)

		vote := &models.ThreadVote{}
		err := unmarshalBody(r, vote)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		thread, err := database.ThreadVote(env, slug, id, vote)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(thread, w)
			return
		}

		switch err.(type) {
		case *models.ErrorNotFound:
			processErrorNotFound(w, err)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
		}
	}
}

func HandleThreadUpdate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, id := parseSlugOrId(r)
		short := &models.ThreadShort{}
		err := unmarshalBody(r, short)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		thread, err := database.ThreadUpdate(
			env,
			slug,
			id,
			short.Title,
			short.Message,
		)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(thread, w)
			return
		}
		switch err.(type) {
		case *models.ErrorNotFound:
			processErrorNotFound(w, err)
		default:
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
