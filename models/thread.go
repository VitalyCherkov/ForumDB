package models

import "time"

//easyjson:json
type ThreadShort struct {
	Author  string     `json:"author"`
	Message string     `json:"message"`
	Title   string     `json:"title"`
	Slug    *string    `json:"slug,omitempty"`
	Created *time.Time `json:"created"`
}

//easyjson:json
type ThreadDetail struct {
	Id    int    `json:"id"`
	Forum string `json:"forum"`
	Votes int    `json:"votes"`
	ThreadShort
}

//easyjson:json
type ThreadDetailList []ThreadDetail
