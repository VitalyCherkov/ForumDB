package database

import (
	"ForumDB/models"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

const (
	queryThreadCreate = `
		INSERT INTO thread (slug, forum, author, title, message, created)
		VALUES (
			$1,
		    (SELECT F.slug FROM forum F WHERE F.slug = $2),
		  	(SELECT U.nickname FROM fuser U WHERE U.nickname = $3),
		    $4,
		    $5,
		    $6
		)
		RETURNING *
	`

	queryThreadGetBySlug = `SELECT * FROM thread WHERE thread.slug = $1`

	queryThreadGetById = `SELECT * FROM thread WHERE thread.id = $1`
)

func ThreadGetBySlugOrId(env *models.Env, slug *string, id *uint64) (thread *models.ThreadDetail, err error) {
	thread = &models.ThreadDetail{}

	if slug != nil {
		err = env.DB.Get(thread, queryThreadGetBySlug, slug)
	} else if id != nil {
		err = env.DB.Get(thread, queryThreadGetById, id)
	} else {
		return nil, &models.ErrorNotFound{
			Message: `can not find thread without neither "slug"" nor "id"`,
		}
	}

	if err == nil {
		return
	}
	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: fmt.Sprintf(
				`thread by request: slug="%v" id="%v" does not exist`,
				slug,
				id,
			),
		}
	}
	return nil, &models.DatabaseError{
		Message: err.Error(),
	}
}

func ThreadCreate(
	env *models.Env,
	short *models.ThreadShort,
	forumSlug string,
) (thread *models.ThreadDetail, err error) {

	thread = &models.ThreadDetail{}
	err = env.DB.Get(
		thread,
		queryThreadCreate,
		short.Slug,
		forumSlug,
		short.Author,
		short.Title,
		short.Message,
		short.Created,
	)

	if err == nil {
		return thread, nil
	}
	fmt.Printf("Thread create error: %s\n", err.Error())

	pqCode := err.(*pq.Error).Code
	if pqCode == uniqueViolationCode {
		fmt.Printf("Thread create conflict %v %v\n", pqCode, uniqueViolationCode)
		thread, err := ThreadGetBySlugOrId(env, short.Slug, nil)
		if err != nil {
			fmt.Printf(err.Error())
			return nil, err
		}
		return nil, &models.ErrorThreadAlreadyExists{
			Thread: thread,
		}
	}
	if pqCode == notNullViolationCode {
		return nil, &models.ErrorNotFound{
			Message: fmt.Sprintf(
				`Thread error: can not create thread by nickname: "%s" and forum: "%s"`,
				short.Author,
				forumSlug,
			),
		}
	}

	return nil, &models.DatabaseError{
		Message: fmt.Sprintf(`Thread error: can not create new thread: %s`, err.Error()),
	}
}

func ThreadGetList(
	env *models.Env,
	forumSlug string,
	since time.Time,
	limit uint64,
	desc bool,
) (threads *models.ThreadDetailList, err error) {
	forumErr := doesForumExist(env, forumSlug)
	if forumErr != nil {
		return nil, forumErr
	}

	q := strings.Builder{}
	q.WriteString("SELECT * FROM thread WHERE forum = $1")

	args := make([]interface{}, 0, 3)
	args = append(args, forumSlug)

	sinceDefault := time.Time{}
	if since != sinceDefault {
		if desc {
			q.WriteString(" AND created <= $2")
		} else {
			q.WriteString(" AND created >= $2")
		}
		args = append(args, since)
	}

	if desc {
		q.WriteString(" ORDER BY created DESC")
	} else {
		q.WriteString(" ORDER BY created")
	}

	if limit > 0 {
		if since != sinceDefault {
			q.WriteString(" LIMIT $3")
		} else {
			q.WriteString(" LIMIT $2")
		}
		args = append(args, limit)
	}

	threads = &models.ThreadDetailList{}
	err = env.DB.Select(threads, q.String(), args...)

	if err == nil {
		return threads, nil
	} else {
		return nil, &models.DatabaseError{
			Message: fmt.Sprintf(`threads list error: %s`, err.Error()),
		}
	}
}
