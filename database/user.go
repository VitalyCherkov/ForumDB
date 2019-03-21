package database

import (
	"ForumDB/models"
	"database/sql"
)

const (
	selectUserByEmailOrNickname = `
		SELECT nickname, email, fullname, about FROM fuser
		WHERE nickname = $1 OR email = $2
	`

	queryInsertUser = `
		INSERT INTO fuser (nickname, email, fullname, about)
		VALUES ($1, $2, $3, $4)
	`
)

func CreateUser(env *models.Env, nickname string, short *models.UserShort) (err error) {
	foundUsers := &models.UserDetailList{}
	err = env.DB.Select(foundUsers, selectUserByEmailOrNickname, nickname, short.Email)

	if err == nil && len(*foundUsers) > 0 {
		return &models.ErrorUserAlreadyExists{
			Users: foundUsers,
		}
	}

	if err == sql.ErrNoRows || len(*foundUsers) == 0 {
		_, err = env.DB.Query(queryInsertUser,
			nickname,
			short.Email,
			short.FullName,
			short.About,
		)

		if err != nil {
			return &models.DatabaseError{
				Message: err.Error(),
			}
		}
		return nil
	}

	return &models.DatabaseError{
		Message: err.Error(),
	}
}
