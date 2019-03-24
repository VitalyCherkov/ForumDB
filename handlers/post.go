package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
)

func HandlePostListCreate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, treadId := parseSlugOrId(r)

		posts := &models.PostDetailList{}
		err := unmarshalBody(r, posts)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		createdPosts, err := database.PostCreateList(env, posts, slug, treadId)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(createdPosts, w)
			return
		}

		switch err.(type) {
		case *models.ErrorNotFound:
			processErrorNotFound(w, err)
		case *models.ErrorPostIncorrectThreadOfParent:
			w.WriteHeader(http.StatusConflict)
			err := err.(*models.ErrorPostIncorrectThreadOfParent)
			err.Message = err.Error()
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err.Error())
		}
	}
}

func HandlePostDetail(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		postCombined := &models.PostCombined{}
		postCombined.Post, err = database.PostGetDetail(env, id)
		if err != nil {
			if !processErrorNotFound(w, err) {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}
			return
		}

		paramsList := strings.Split(
			r.URL.Query().Get("related"),
			",",
		)
		for _, param := range paramsList {
			switch param {
			case "thread":
				postCombined.Thread, err = database.ThreadGetBySlugOrId(
					env,
					nil,
					&postCombined.Post.Thread,
				)
			case "user":
				postCombined.Author, err = database.UserGet(
					env,
					postCombined.Post.Author,
				)
			case "forum":
				postCombined.Forum, err = database.ForumGet(
					env,
					postCombined.Post.Forum,
				)
			}
			if err != nil {
				if !processErrorNotFound(w, err) {
					fmt.Println(err.Error())
					w.WriteHeader(http.StatusBadRequest)
				}
				return
			}
		}

		_, _, _ = easyjson.MarshalToHTTPResponseWriter(postCombined, w)
	}
}
