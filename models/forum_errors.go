package models

import "fmt"

type ErrorForumAlreadyExists struct {
	Forum *ForumDetail
}

func (err *ErrorForumAlreadyExists) Error() string {
	return fmt.Sprintf(`Forum error: forum with slug "%s" already exists`, err.Forum.Slug)
}
