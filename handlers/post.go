package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func HandlePostListCreate(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug, treadId := parseSlugOrId(mux.Vars(r)["slug_or_id"])

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
			err := err.(*models.ErrorNotFound)
			w.WriteHeader(http.StatusNotFound)
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
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
