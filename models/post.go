package models

import "time"

//easyjson:json
type PostDetail struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Parent  uint64 `json:"parent"`

	Id       uint64    `json:"id,omitempty"`
	Created  time.Time `json:"created,omitempty"`
	Forum    string    `json:"forum,omitempty"`
	Thread   uint64    `json:"thread,omitempty"`
	IsEdited bool      `json:"isEdited,omitempty"`

	Path []int64 `json:"-"`
}

//easyjson:json
type PostDetailList []PostDetail

//easyjson:json
type PostCombined struct {
	Post   *PostDetail   `json:"post"`
	Author *UserDetail   `json:"author,omitempty"`
	Thread *ThreadDetail `json:"thread,omitempty"`
	Forum  *ForumDetail  `json:"forum,omitempty"`
}
