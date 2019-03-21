package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"

	"github.com/mailru/easyjson"

	"github.com/gorilla/mux"
)

func HandleUserCreate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nickname := vars["nickname"]

		user := &models.UserShort{}
		err := unmarshalBody(r, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = database.UserCreate(env, nickname, user)

		if err == nil {
			w.WriteHeader(http.StatusCreated)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(&models.UserDetail{
				Nickname:  nickname,
				UserShort: *user,
			}, w)
			return
		}

		switch err.(type) {
		case *models.ErrorUserAlreadyExists:
			err := err.(*models.ErrorUserAlreadyExists)
			w.WriteHeader(http.StatusConflict)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err.Users, w)
			return
		case *models.DatabaseError:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func HandleUserGet(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nickname := vars["nickname"]

		user, err := database.UserGet(env, nickname)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(user, w)
			return
		}
		switch err.(type) {
		case *models.ErrorNotFound:
			err := err.(*models.ErrorNotFound)
			w.WriteHeader(http.StatusNotFound)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandleUserUpdate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nickname := vars["nickname"]

		user := &models.UserShort{}
		err := unmarshalBody(r, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = database.UserUpdate(env, nickname, user)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(&models.UserDetail{
				Nickname:  nickname,
				UserShort: *user,
			}, w)
			return
		}
		switch err.(type) {
		case *models.ErrorConflict:
			w.WriteHeader(http.StatusConflict)
			err := err.(*models.ErrorConflict)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
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
