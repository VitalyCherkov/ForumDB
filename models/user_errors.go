package models

import "fmt"

type ErrorUserAlreadyExists struct {
	Users *UserDetailList
}

func (err *ErrorUserAlreadyExists) Error() string {
	return fmt.Sprintf("User error: user with this email or nickname already exists: %v", *err.Users)
}
