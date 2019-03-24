package database

import (
	"ForumDB/models"
	"database/sql"
	"strconv"
	"strings"

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

func userIsNoUpdate(user *models.UserDetail) bool {
	return user.Nickname == "" && user.Email == "" && user.About == "" && user.FullName == ""
}

func userBuildUpdateQuery(nickname string, user *models.UserDetail) (string, []interface{}) {
	q := strings.Builder{}
	q.WriteString("UPDATE fuser SET")
	args := make([]interface{}, 0, 5)

	buildItem := func(fieldValue, fieldName string) {
		if fieldValue != "" {
			args = append(args, fieldValue)
			count := len(args)
			if count > 1 {
				q.WriteString(",")
			}
			q.WriteString(" " + fieldName + " = $" + strconv.Itoa(count))
		}
	}

	buildItem(user.Nickname, "nickname")
	buildItem(user.Email, "email")
	buildItem(user.FullName, "fullname")
	buildItem(user.About, "about")
	args = append(args, nickname)
	q.WriteString(" WHERE fuser.nickname = $" + strconv.Itoa(len(args)) + " RETURNING *")

	return q.String(), args
}

func UserUpdate(env *models.Env, nickname string, detail *models.UserDetail) (user *models.UserDetail, err error) {
	if userIsNoUpdate(detail) {
		return UserGet(env, nickname)
	}

	query, args := userBuildUpdateQuery(nickname, detail)
	user = &models.UserDetail{}
	err = env.DB.Get(user, query, args...)
	if err == nil {
		return user, nil
	}

	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: "Can not find user with such nickname to update",
		}
	}
	if err.(*pq.Error).Code == uniqueViolationCode {
		return nil, &models.ErrorConflict{
			Message: err.Error(),
		}
	}
	return nil, &models.DatabaseError{
		Message: err.Error(),
	}
}
