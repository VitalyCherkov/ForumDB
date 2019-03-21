package models

import "fmt"

type DatabaseError struct {
	Message string
}

func (err *DatabaseError) Error() string {
	return fmt.Sprintf("Database error: %s", err.Message)
}
