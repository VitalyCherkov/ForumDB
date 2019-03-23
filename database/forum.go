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
