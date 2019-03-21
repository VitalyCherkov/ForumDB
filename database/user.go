package database

import (
	"ForumDB/models"
	"database/sql"

	"github.com/lib/pq"
)

const (
	querySelectUserByEmailOrNickname = `
		SELECT nickname, email, fullname, about FROM fuser
		WHERE nickname = $1 OR email = $2
	`

	querySelectUserByNickname = `
		SELECT nickname, email, fullname, about FROM fuser
		WHERE nickname = $1
	`

	queryInsertUser = `
		INSERT INTO fuser (nickname, email, fullname, about)
		VALUES ($1, $2, $3, $4)
	`

	queryUpdateUser = `
		UPDATE fuser
		SET email = $2,
		    fullname = $3,
		    about = $4
		WHERE nickname = $1
		RETURNING *
	`
)

func UserCreate(env *models.Env, nickname string, short *models.UserShort) (err error) {
	foundUsers := &models.UserDetailList{}
	err = env.DB.Select(foundUsers, querySelectUserByEmailOrNickname, nickname, short.Email)

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

func UserGet(env *models.Env, nickname string) (user *models.UserDetail, err error) {
	user = &models.UserDetail{}
	err = env.DB.Get(user, querySelectUserByNickname, nickname)
	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: "Can not find user with nickname",
		}
	}
	if err != nil {
		return nil, &models.DatabaseError{
			Message: err.Error(),
		}
	}
	return user, nil
}

func UserUpdate(env *models.Env, nickname string, short *models.UserShort) (err error) {
	_, err = env.DB.Query(queryUpdateUser, nickname, short.Email, short.FullName, short.About)
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return &models.ErrorNotFound{
			Message: "Can not find user with such nickname to update",
		}
	}
	if err.(*pq.Error).Code == uniqueViolationCode {
		return &models.ErrorConflict{
			Message: err.Error(),
		}
	}
	return &models.DatabaseError{
		Message: err.Error(),
	}
}
