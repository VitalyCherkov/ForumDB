package database

import (
	"ForumDB/models"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

const (
	colsToInsert    = 7
	lastRowTemplate = "($%d, $%d, $%d, $%d, $%d, $%d, $%d)"
	rowTemplate     = lastRowTemplate + ","

	queryPostGetDetail = `SELECT * FROM post WHERE post.id = $1`
)

func PostCreateList(
	env *models.Env,
	posts *models.PostDetailList,
	slug *string,
	threadId *uint64,
) (createdPosts *models.PostDetailList, err error) {
	targetThread, err := ThreadGetBySlugOrId(env, slug, threadId)
	if err != nil {
		return nil, err
	}
	postCount := len(*posts)
	if postCount == 0 {
		return &models.PostDetailList{}, nil
	}

	usersMap := make(map[string]string)
	parentPostsMap := make(map[uint64]interface{})
	for index, p := range *posts {
		lower := strings.ToLower(p.Author)
		(*posts)[index].Author = lower
		usersMap[lower] = p.Author
		if p.Parent != 0 {
			parentPostsMap[p.Parent] = struct{}{}
		}
	}

	transaction, err := env.DB.Beginx()
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "post create error: " + err.Error(),
		}
	}
	defer transaction.Rollback()

	// Проверка сущестования всех юзеров
	userStatement, err := transaction.Preparex(`SELECT nickname FROM fuser WHERE nickname = $1`)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "post create error: " + err.Error(),
		}
	}
	defer userStatement.Close()
	for k := range usersMap {
		nickname := ""
		err := userStatement.Get(&nickname, k)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, &models.ErrorNotFound{
					Message: fmt.Sprintf(
						`posts create: can not find user with nickname = "%s"`,
						k,
					),
				}
			} else {
				return nil, &models.DatabaseError{
					Message: "posts create: " + err.Error(),
				}
			}
		} else {
			usersMap[k] = nickname
		}
	}

	// Проверка существования всех родительских постов
	parentPostStatement, err := transaction.Preparex(`SELECT thread FROM post WHERE post.id = $1`)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "posts create: " + err.Error(),
		}
	}
	defer parentPostStatement.Close()
	for k := range parentPostsMap {
		parentThreadId := uint64(0)
		err = parentPostStatement.QueryRow(k).Scan(&parentThreadId)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, &models.ErrorPostIncorrectThreadOfParent{
					Message: fmt.Sprintf(`post create: can not find parent post with id = "%d"`, k),
				}
			} else {
				return nil, &models.DatabaseError{
					Message: "posts create: " + err.Error(),
				}
			}
		} else if parentThreadId != targetThread.Id {
			return nil, &models.ErrorPostIncorrectThreadOfParent{
				CurThreadId:    targetThread.Id,
				ParentThreadId: parentThreadId,
			}
		}
	}

	// Выборка всех будущих id
	idsStatement, err := transaction.Preparex(`SELECT nextval(pg_get_serial_sequence('post', 'id'))`)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "post create: unable to prepare stmt for ids: " + err.Error(),
		}
	}
	defer idsStatement.Close()
	for index := range *posts {
		err = idsStatement.QueryRow().Scan(&(*posts)[index].Id)
		if err != nil {
			return nil, &models.DatabaseError{
				Message: fmt.Sprintf(`post create: unable to select new id from series`),
			}
		}
	}

	// Составление запроса для вставки новых записей
	query := strings.Builder{}
	now := time.Time{}
	err = env.DB.QueryRow(`SELECT * FROM now()`).Scan(&now)
	if err != nil {
		return nil, err
	}

	query.WriteString("INSERT INTO post (id, author, thread, forum, message, parent, created) VALUES ")
	args := make([]interface{}, 0, colsToInsert*postCount)

	rowArgIndexes := []interface{}{1, 2, 3, 4, 5, 6, 7}
	for index, post := range *posts {
		args = append(
			args,
			(*posts)[index].Id,
			usersMap[post.Author],
			targetThread.Id,
			targetThread.Forum,
			post.Message,
			post.Parent,
			now,
		)
		(*posts)[index].Author = usersMap[post.Author]
		(*posts)[index].Thread = targetThread.Id
		(*posts)[index].Forum = targetThread.Forum
		(*posts)[index].Created = now

		if index+1 == postCount {
			query.WriteString(fmt.Sprintf(lastRowTemplate, rowArgIndexes...))
		} else {
			query.WriteString(fmt.Sprintf(rowTemplate, rowArgIndexes...))
		}
		for i, rowIndex := range rowArgIndexes {
			rowArgIndexes[i] = rowIndex.(int) + colsToInsert
		}
	}
	_, err = transaction.Exec(query.String(), args...)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "can not inset posts: " + err.Error(),
		}
	}

	// Обновление количества постов в формуме
	_, err = transaction.Exec(
		`UPDATE forum SET posts = posts + $1 WHERE forum.slug = $2`,
		&postCount,
		&targetThread.Forum,
	)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "can not update post count in forum: " + err.Error(),
		}
	}

	err = transaction.Commit()
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "can not commit transaction: " + err.Error(),
		}
	}

	return posts, nil
}

func PostGetDetail(env *models.Env, id uint64) (post *models.PostDetail, err error) {
	post = &models.PostDetail{}
	err = env.DB.Get(post, queryPostGetDetail, &id)
	if err == nil {
		return post, nil
	}
	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: fmt.Sprintf("post detail: can not find post by id=%d", id),
		}
	}
	return nil, &models.DatabaseError{
		Message: "post detail: " + err.Error(),
	}
}
