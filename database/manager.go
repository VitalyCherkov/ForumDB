package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(user, password, port, dbName string) (db *sqlx.DB, err error) {

	connectionString := fmt.Sprintf(
		"postgresql://%v:%v@localhost:%v/%v?sslmode=disable",
		user,
		password,
		port,
		dbName,
	)

	db, err = sqlx.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Printf("successfully connected to database")
	return db, nil
}
