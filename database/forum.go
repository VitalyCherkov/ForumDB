package database

import (
	"ForumDB/models"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

const (
	queryForumGet = `
		SELECT * FROM forum WHERE forum.slug = $1
	`

	queryForumCreate = `
		INSERT INTO forum (slug, author, title)
		VALUES ($1, (SELECT nickname FROM fuser WHERE nickname = $2), $3)
		RETURNING *
	`

	queryForumCheckBySlug = `SELECT slug FROM forum WHERE slug = $1`
)

func ForumGet(env *models.Env, slug string) (forum *models.ForumDetail, err error) {
	forum = &models.ForumDetail{}
	err = env.DB.Get(forum, queryForumGet, slug)

	if err == nil {
		return
	}
	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: slug,
		}
	}

	return nil, &models.DatabaseError{
		Message: fmt.Sprintf(
			`can not find forum with slug: "%s". %s`,
			slug,
			err.Error(),
		),
	}
}

func ForumCreate(env *models.Env, short *models.ForumSort) (forum *models.ForumDetail, err error) {
	forum = &models.ForumDetail{}
	err = env.DB.Get(forum, queryForumCreate, short.Slug, short.Author, short.Title)
	if err == nil {
		return forum, nil
	}
	pqCode := err.(*pq.Error).Code
	if pqCode == uniqueViolationCode {
		fmt.Printf("Forum create conflict %v %v\n", pqCode, uniqueViolationCode)
		forum, err := ForumGet(env, short.Slug)
		if err != nil {
			fmt.Printf("Forum get error, %v", err.Error())
			return nil, err
		}
		return nil, &models.ErrorForumAlreadyExists{
			Forum: forum,
		}
	}
	if pqCode == notNullViolationCode {
		return nil, &models.ErrorNotFound{
			Message: fmt.Sprintf(
				`Forum error: can not find user by nickname: "%s"`,
				short.Author,
			),
		}
	}

	return nil, &models.DatabaseError{
		Message: fmt.Sprintf(`Forum error: can not create new forum: %s`, err.Error()),
	}
}

func doesForumExist(env *models.Env, slug string) (err error) {
	foundSlug := ""
	err = env.DB.Get(&foundSlug, queryForumCheckBySlug, slug)
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return &models.ErrorNotFound{
			Message: fmt.Sprintf(`Forum error: can not find with slug %s`, slug),
		}
	} else {
		return &models.DatabaseError{
			Message: fmt.Sprintf(`Forum error: %s`, err.Error()),
		}
	}
}

func ForumGetUsers(
	env *models.Env,
	slug, since string,
	limit uint64,
	desc bool,
) (users *models.UserDetailList, err error) {

	err = doesForumExist(env, slug)
	if err != nil {
		return nil, err
	}

	args := make([]interface{}, 0, 3)
	args = append(args, slug)

	var nicknameCmpPart string
	if since != "" {
		if desc {
			nicknameCmpPart = "AND F.nickname < $2"
		} else {
			nicknameCmpPart = "AND F.nickname > $2"
		}
		args = append(args, since)
	}
	var limitPart string
	if limit != 0 {
		if since != "" {
			limitPart = " LIMIT $3"
		} else {
			limitPart = " LIMIT $2"
		}
		args = append(args, limit)
	}
	var descPart string
	if desc {
		descPart = " DESC"
	} else {
		descPart = " ASC"
	}

	query := fmt.Sprintf(`
		SELECT F.nickname as nickname, fullname, email, about FROM
			forum_fuser F
			JOIN fuser ON F.nickname = fuser.nickname 
			WHERE slug = $1 %s
			ORDER BY F.nickname %s %s
	`, nicknameCmpPart, descPart, limitPart)

	users = &models.UserDetailList{}
	err = env.DB.Select(users, query, args...)
	if err == nil {
		return users, nil
	}
	if err.(*pq.Error).Code == notNullViolationCode {
		return nil, &models.ErrorNotFound{
			Message: `can not get forum users by slug: ` + slug,
		}
	}
	return nil, &models.DatabaseError{
		Message: `can not get forum users: ` + err.Error(),
	}
}
