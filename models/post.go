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
	Thread   int       `json:"thread,omitempty"`
	IsEdited string    `json:"isEdited,omitempty"`
}

//easyjson:json
type PostDetailList []PostDetail
