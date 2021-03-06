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
		case *models.ErrorPostIncorrectThreadOfParent:
			w.WriteHeader(http.StatusConflict)
			err := err.(*models.ErrorPostIncorrectThreadOfParent)
			err.Message = err.Error()
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
		default:
			if !processErrorNotFound(w, err) {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
			}
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

func HandlePostUpdate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		post := &models.PostDetail{}
		err = unmarshalBody(r, post)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		post, err = database.PostUpdate(env, id, post)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(post, w)
			return
		}
		if !processErrorNotFound(w, err) {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func HandlePostListGet(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, id := parseSlugOrId(r)

		sortType := r.URL.Query().Get("sort")
		limit, err := parseUint64FromQuery(r, "limit")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		since, err := parseUint64FromQuery(r, "since")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		desc := r.URL.Query().Get("desc") == "true"

		posts, err := database.PostListGet(env, slug, id, sortType, since, limit, desc)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(posts, w)
			return
		}
		if !processErrorNotFound(w, err) {
			fmt.Println("Post list get: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
