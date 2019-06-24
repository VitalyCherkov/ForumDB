package database

import "ForumDB/models"

func ServiceClear(env *models.Env) {
	_, _ = env.DB.Exec(`TRUNCATE TABLE fuser CASCADE`)
	_, _ = env.DB.Exec(`TRUNCATE TABLE post CASCADE`)
	_, _ = env.DB.Exec(`TRUNCATE TABLE forum_fuser CASCADE`)
	_, _ = env.DB.Exec(`TRUNCATE TABLE forum CASCADE`)
}

func ServiceStatus(env *models.Env) (status *models.ServiceStatus, err error) {
	status = &models.ServiceStatus{}
	err = env.DB.QueryRow(`SELECT COUNT(*) FROM post`).Scan(&(*status).Post)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "service status: can not get post count" + err.Error(),
		}
	}

	err = env.DB.QueryRow(`SELECT COUNT(*) FROM thread`).Scan(&(*status).Thread)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "service status: can not get thread count" + err.Error(),
		}
	}

	err = env.DB.QueryRow(`SELECT COUNT(*) FROM forum`).Scan(&(*status).Forum)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "service status: can not get forum count" + err.Error(),
		}
	}

	err = env.DB.QueryRow(`SELECT COUNT(*) FROM fuser`).Scan(&(*status).User)
	if err != nil {
		return nil, &models.DatabaseError{
			Message: "service status: can not get user count" + err.Error(),
		}
	}

	return status, nil
}
