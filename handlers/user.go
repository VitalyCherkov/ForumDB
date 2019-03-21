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

		err = database.CreateUser(env, nickname, user)

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
			fmt.Println(err.Error())
			return
		}

		fmt.Printf(nickname, *user)
	}
}

func HandleUserGet(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nickname := vars["nickname"]

		fmt.Printf(nickname)
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

		fmt.Printf(nickname, *user)
	}
}
