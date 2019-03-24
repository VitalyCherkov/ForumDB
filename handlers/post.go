package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"

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
