package models

import (
	"fmt"
	"net/http"
)

type DatabaseError struct {
	Message string
}

func (err *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %s", err.Message)
}

//easyjson:json
type ErrorNotFound struct {
	Message string `json:"message"`
}

func (err *ErrorNotFound) Header() http.Header {
	panic("implement me")
}

func (err *ErrorNotFound) Write([]byte) (int, error) {
	panic("implement me")
}

func (err *ErrorNotFound) WriteHeader(statusCode int) {
	panic("implement me")
}

func (err *ErrorNotFound) Error() string {
	return fmt.Sprintf("Not found in database error: %s", err.Message)
}

//easyjson:json
type ErrorConflict struct {
	Message string `json:"message"`
}

func (err *ErrorConflict) Error() string {
	return fmt.Sprintf("Conflict error: %s", err.Message)
}
