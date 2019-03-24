package models

import "fmt"

//easyjson:json
type ErrorPostIncorrectThreadOfParent struct {
	CurThreadId    uint64
	ParentThreadId uint64
	Message        string `json:"message"`
}

func (err *ErrorPostIncorrectThreadOfParent) Error() string {
	return fmt.Sprintf(
		`Post error: incorrect parent post threadId="%d", cur="%d"`,
		err.ParentThreadId,
		err.CurThreadId,
	)
}
