package database

import (
	"ForumDB/models"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

const (
	colsToInsert    = 8
	lastRowTemplate = "($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)"
	rowTemplate     = lastRowTemplate + ","

	queryPostGetDetail = `
		SELECT id, author, created, forum, isEdited, message, parent, thread
		FROM post WHERE post.id = $1
	`

	queryPostUpdate = `
		UPDATE post SET message = $1 WHERE post.id = $2
        RETURNING id, author, created, forum, isEdited, message, parent, thread
	`
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
	parentPostsMap := make(map[uint64][]int64)
	for index, p := range *posts {
		lower := strings.ToLower(p.Author)
		(*posts)[index].Author = lower
		usersMap[lower] = p.Author
		if p.Parent != 0 {
			parentPostsMap[p.Parent] = nil
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
					Message: "posts create: users check: " + err.Error(),
				}
			}
		} else {
			usersMap[k] = nickname
		}
	}

	// Проверка существования всех родительских постов
	parentPostStatement, err := transaction.Preparex(`SELECT thread, path FROM post WHERE post.id = $1`)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "posts create: " + err.Error(),
		}
	}
	defer parentPostStatement.Close()
	for k := range parentPostsMap {
		parentThreadId := uint64(0)
		var path []int64
		err = parentPostStatement.QueryRow(k).Scan(&parentThreadId, pq.Array(&path))
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
		parentPostsMap[k] = path
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

	query.WriteString("INSERT INTO post (id, author, thread, forum, message, parent, created, path) VALUES ")
	args := make([]interface{}, 0, colsToInsert*postCount)

	rowArgIndexes := []interface{}{1, 2, 3, 4, 5, 6, 7, 8}
	for index, post := range *posts {
		args = append(
			args,
			post.Id,
			usersMap[post.Author],
			targetThread.Id,
			targetThread.Forum,
			post.Message,
			post.Parent,
			now,
			pq.Array(append(parentPostsMap[post.Parent], int64(post.Id))),
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

	// Добавление пользователей форума в таблицу forum_fuser
	query = strings.Builder{}
	userCount := len(usersMap)
	args = make([]interface{}, 0, userCount+1)
	args = append(args, targetThread.Forum)
	query.WriteString("INSERT INTO forum_fuser(slug, nickname) VALUES")
	userCurIndex := 1
	for _, nickname := range usersMap {
		if userCurIndex != userCount {
			query.WriteString(" ($1, $" + strconv.Itoa(userCurIndex+1) + "),")
		} else {
			query.WriteString(" ($1, $" + strconv.Itoa(userCurIndex+1) + ")")
		}
		args = append(args, nickname)
	}
	query.WriteString(` ON CONFLICT DO NOTHING`)
	_, err = transaction.Exec(query.String(), args...)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "can not insert users into forum_fuser: " + err.Error(),
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

func PostUpdate(env *models.Env, id uint64, short *models.PostDetail) (post *models.PostDetail, err error) {
	if short.Message == "" {
		return PostGetDetail(env, id)
	}
	post = &models.PostDetail{}
	err = env.DB.Get(post, queryPostUpdate, short.Message, id)
	if err == nil {
		return post, nil
	}
	if err == sql.ErrNoRows {
		return nil, &models.ErrorNotFound{
			Message: fmt.Sprintf(`post update: can not find post with id="%d"`, id),
		}
	}
	return nil, &models.DatabaseError{
		Message: "post update: " + err.Error(),
	}
}

func PostListGet(
	env *models.Env,
	slug *string,
	threadId *uint64,
	sortType string,
	since, limit uint64,
	desc bool,
) (posts *models.PostDetailList, err error) {
	targetThread, err := ThreadGetBySlugOrId(env, slug, threadId)
	if err != nil {
		return nil, err
	}

	query := strings.Builder{}
	query.WriteString(`
		SELECT P.id, P.author, P.created, P.forum, P.isEdited, P.message, P.parent, P.thread
		FROM post P
	`)
	args := make([]interface{}, 0, 3)
	args = append(args, targetThread.Id)

	switch sortType {
	case "tree":
		if since > 0 {
			query.WriteString(` JOIN post Pfilter ON Pfilter.id = $2`)
			if desc {
				query.WriteString(` WHERE P.thread = $1 AND P.path < Pfilter.path`)
			} else {
				query.WriteString(` WHERE P.thread = $1 AND P.path > Pfilter.path`)
			}
			args = append(args, since)
		} else {
			query.WriteString(` WHERE P.thread = $1`)
		}

		if desc {
			query.WriteString(` ORDER BY P.path DESC`)
		} else {
			query.WriteString(` ORDER BY P.path`)
		}

		if limit > 0 {
			if since > 0 {
				query.WriteString(` LIMIT $3`)
			} else {
				query.WriteString(` LIMIT $2`)
			}
			args = append(args, limit)
		}

	case "parent_tree":
		// Оптимизировать запрос
		query.WriteString(` WHERE P.path[1] IN (SELECT Pfilter.id FROM post Pfilter`)
		if since > 0 {
			query.WriteString(` JOIN post Pedge ON Pedge.id = $2`)
			if desc {
				query.WriteString(` WHERE Pfilter.thread = $1 AND Pfilter.parent = 0 AND Pfilter.id < Pedge.path[1]`)
			} else {
				query.WriteString(` WHERE Pfilter.thread = $1 AND Pfilter.parent = 0 AND Pfilter.id > Pedge.path[1]`)
			}
			args = append(args, since)
		} else {
			query.WriteString(` WHERE Pfilter.thread = $1 AND Pfilter.parent = 0`)
		}
		if desc {
			query.WriteString(` ORDER BY Pfilter.id DESC`)
		} else {
			query.WriteString(` ORDER BY Pfilter.id`)
		}
		if limit > 0 {
			if since > 0 {
				query.WriteString(` LIMIT $3`)
			} else {
				query.WriteString(` LIMIT $2`)
			}
			args = append(args, limit)
		}
		if desc {
			query.WriteString(` ) ORDER BY P.path[1] DESC, path`)
		} else {
			query.WriteString(` ) ORDER BY P.path`)
		}
	default:
		// flat
		query.WriteString(` WHERE P.thread = $1`)
		if since > 0 {
			if desc {
				query.WriteString(` AND P.id < $2`)
			} else {
				query.WriteString(` AND P.id > $2`)
			}
			args = append(args, since)
		}
		if desc {
			query.WriteString(` ORDER BY P.id DESC`)
		} else {
			query.WriteString(` ORDER BY P.id ASC`)
		}
		if limit > 0 {
			if since > 0 {
				query.WriteString(` LIMIT $3`)
			} else {
				query.WriteString(` LIMIT $2`)
			}
			args = append(args, limit)
		}
	}

	posts = &models.PostDetailList{}
	err = env.DB.Select(posts, query.String(), args...)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: fmt.Sprintf(
				`post list get: sort_type="%s", limit="%d", since="%d", desc="%v"\n\terr: %s`,
				sortType,
				limit,
				since,
				desc,
				err.Error(),
			),
		}
	}
	return posts, nil
}
