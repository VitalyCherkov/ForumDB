package handlers

import (
	"ForumDB/models"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"

	"github.com/gorilla/mux"
)

func parseSlugOrId(r *http.Request) (slug *string, id *uint64) {
	slugOrId := mux.Vars(r)["slug_or_id"]
	_id, err := strconv.ParseUint(slugOrId, 10, 64)
	if err == nil {
		return nil, &_id
	} else {
		return &slugOrId, nil
	}
}

func processErrorNotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	if err, ok := err.(*models.ErrorNotFound); ok {
		_, _, _ = easyjson.MarshalToHTTPResponseWriter(err, w)
	}
}
