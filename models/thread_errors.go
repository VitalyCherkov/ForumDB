package models

import "fmt"

type ErrorThreadAlreadyExists struct {
	Thread *ThreadDetail
}

func (err *ErrorThreadAlreadyExists) Error() string {
	return fmt.Sprintf(`Thread error: thread with slug "%v" already exists`, err.Thread.Slug)
}
