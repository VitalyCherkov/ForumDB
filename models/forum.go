package models

//easyjson:json
type ForumSort struct {
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	Author string `json:"user"`
}

//easyjson:json
type ForumDetail struct {
	ForumSort

	Posts   int `json:"posts"`
	Threads int `json:"threads"`
}
