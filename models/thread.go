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
	Id    uint64 `json:"id"`
	Forum string `json:"forum"`
	Votes int    `json:"votes"`
	ThreadShort
}

//easyjson:json
type ThreadDetailList []ThreadDetail

//easyjson:json
type ThreadVote struct {
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}
