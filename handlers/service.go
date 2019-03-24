package handlers

import (
	"ForumDB/database"
	"ForumDB/models"
	"fmt"
	"net/http"

	"github.com/mailru/easyjson"
)

func HandleServiceClear(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		database.ServiceClear(env)
	}
}

func HandleServiceStatus(env *models.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, err := database.ServiceStatus(env)
		if err == nil {
			_, _, _ = easyjson.MarshalToHTTPResponseWriter(status, w)
		} else {
			fmt.Println("service status error: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
